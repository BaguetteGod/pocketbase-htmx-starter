package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pb-starter/app/components/inputs"
	"pb-starter/app/components/toast"
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

func handleConfirmEmailChangeRequest(form forms.ConfirmEmailChangeForm) *http.Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("%s/api/collections/users/confirm-email-change", os.Getenv("APP_URL"))

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

	if resp.StatusCode != 204 {
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
		return lib.Render(c, profile.ProfilePage(c, form, "", user.Get("oauth") == true))
	})

	group.POST("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, profile.ProfilePage(c, form, string(formErrors), user.Get("oauth") == true))
		}

		message := ""
		if user.Get("oauth") != true && form.Email != user.Email() {
			handleRequestEmailChange(c, form.Email)
			message = fmt.Sprintf("An email has been sent to %s. Please follow the instructions to confirm your new email.", form.Email)
		}

		if err := user.SetUsername(form.Username); err != nil {
			return err
		}

		if err := app.Dao().SaveRecord(user); err != nil {
			return err
		}

		toast.New(c, toast.Toast{Message: message})
		return lib.Render(c, profile.ProfilePage(c, form, "", user.Get("oauth") == true))
	})

	group.POST("/profile/username", func(c echo.Context) error {
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/profile/username", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/profile/username"}.Comp())
	})

	group.POST("/profile/email", func(c echo.Context) error {
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/profile/email", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/profile/email"}.Comp())
	})

	group.GET("/confirm-email-change/:token", func(c echo.Context) error {
		token := c.PathParam("token")
		if len(token) < 200 {
			return c.Redirect(302, "/dashboard")
		}
		form := forms.ConfirmEmailChangeForm{Token: token}

		return lib.Render(c, profile.ConfirmEmailChangePage(form, ""))
	})

	group.POST("/confirm-email-change/:token", func(c echo.Context) error {
		form := forms.GetConfirmEmailChangeForm(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, profile.ConfirmEmailChangePage(form, string(formErrors)))
		}

		resp := handleConfirmEmailChangeRequest(form)
		if resp != nil {
			toast.New(c, toast.Toast{Level: toast.DANGER, Title: "Link is invalid or has expired"})
			return lib.Render(c, profile.ConfirmEmailChangePage(form, ""))
		}

		return lib.Render(c, profile.EmailChangedSuccessPage())
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

		src, err := filesystem.NewFileFromMultipart(file)
		if !lib.ValidateImage(src) {
			toast.New(c, toast.Toast{Level: toast.DANGER, Title: "Invalid image format or size", Message: "Accepted image formats are JPG, WEBP, GIF or PNG, 5MB max."})
			return c.JSON(400, http.StatusBadRequest)
		}

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

		toast.New(c, toast.Toast{})
		return lib.Render(c, profile.ProfilePageContent(c, pfForm, "", user.Get("oauth") == true))
	})
}
