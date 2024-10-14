package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	middleware "pb-starter/app/middelware"
	"pb-starter/app/views/profile"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	pbforms "github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/tools/filesystem"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func handleRequestEmailChange(c echo.Context, email string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenCookie, err := c.Cookie(middleware.AuthCookieName)
	if err != nil {
		fmt.Println("Error getting cookie:", err)
		return
	}

	url := fmt.Sprintf("%s/api/collections/users/request-email-change", os.Getenv("APP_URL"))

	jsonDataMap := map[string]interface{}{
		"newEmail": email,
	}
	jsonData, err := json.Marshal(jsonDataMap)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", tokenCookie.Value)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}

func handleConfirmEmailChangeRequest(form forms.ConfirmPasswordResetFormValue) *http.Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("%s/api/collections/users/confirm-password-reset", os.Getenv("APP_URL"))

	jsonData, err := json.Marshal(form)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.Status != "204" {
		return resp
	}

	return nil
}

func RegisterProfileRoutes(e *core.ServeEvent, group echo.Group, app *pocketbase.PocketBase) {
	group.GET("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.ProfileFormValue{
			Username: user.Username(),
			Email:    user.Email(),
		}
		return lib.Render(c, profile.Profile(c, form, "", user.Get("oauth") == true))
	})

	group.POST("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, profile.Profile(c, form, string(formErrors), user.Get("oauth") == true))
		}

		if user.Get("oauth") != true && form.Email != user.Email() {
			handleRequestEmailChange(c, form.Email)
		}

		if err := user.SetUsername(form.Username); err != nil {
			return err
		}

		if err := app.Dao().SaveRecord(user); err != nil {
			return err
		}

		return lib.Render(c, profile.Profile(c, form, "", user.Get("oauth") == true))
	})

	group.POST("/avatar", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		file, err := c.FormFile("avatar")
		pfForm := forms.ProfileFormValue{
			Username: user.Username(),
			Email:    user.Email(),
		}

		if err != nil {
			return err
		}

		form := pbforms.NewRecordUpsert(app, user)
		if err := form.LoadData(map[string]any{
			"username": user.Username(),
			"email":    user.Email(),
		}); err != nil {
			return err
		}

		src, err := filesystem.NewFileFromMultipart(file)
		if err != nil {
			app.Logger().Error("failed to create file", "error", err)
			return err
		}

		if err := form.AddFiles("avatar", src); err != nil {
			app.Logger().Error(err.Error())
			return err
		}

		if err := form.Submit(); err != nil {
			app.Logger().Error(err.Error())
			return err
		}
		return lib.Render(c, profile.Profile(c, pfForm, "", user.Get("oauth") == true))
	})
}
