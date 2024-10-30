package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pb-starter/app/components/toast"
	"pb-starter/app/controllers/validation_routes"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	middleware "pb-starter/app/middelware"
	"pb-starter/app/views/profile"

	"github.com/fatih/structs"
	"github.com/pocketbase/pocketbase"
	pbforms "github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterProfileRoutes(e *core.ServeEvent, group echo.Group, app *pocketbase.PocketBase) {
	group.GET("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.ProfileFormValue{
			Username:    user.Username(),
			NewEmail:    user.Email(),
			DisplayName: user.GetString("displayname"),
		}
		return lib.Render(c, profile.ProfilePage(c, form, "", user.GetBool("oauth")))
	})

	group.POST("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, profile.ProfilePage(c, form, string(formErrors), user.GetBool("oauth")))
		}

		message := ""
		resp, _ := handleRequestEmailChange(c, form, user)
		if resp != nil && resp.StatusCode == 204 {
			message = fmt.Sprintf("An email has been sent to %s. Please follow the instructions to confirm your new email.", form.NewEmail)
		}

		if err := user.SetUsername(form.Username); err != nil {
			return err
		}

		user.Set("displayname", form.DisplayName)

		if err := app.Dao().SaveRecord(user); err != nil {
			return err
		}

		toast.New(c, toast.Toast{Message: message})
		return lib.Render(c, profile.ProfilePage(c, form, "", user.GetBool("oauth")))
	})

	group.GET("/confirm-email-change/:token", func(c echo.Context) error {
		token := c.PathParam("token")
		user, _ := e.App.Dao().FindAuthRecordByToken(token, e.App.Settings().RecordEmailChangeToken.Secret)
		if user == nil {
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

		data := structs.Map(form)
		r := lib.PocketBaseRequest{
			Route:  "/users/confirm-email-change",
			Method: "POST",
			Data:   data,
		}
		resp, _ := lib.NewPocketBaseRequest(r)
		if resp.StatusCode != 204 {
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
			NewEmail: user.Email(),
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
		return lib.Render(c, profile.ProfilePageContent(c, pfForm, "", user.GetBool("oauth")))
	})

	validation_routes.RegisterProfileValidationRoutes(e, group)
}

func handleRequestEmailChange(c echo.Context, form forms.ProfileFormValue, user *models.Record) (*http.Response, error) {
	if user.GetBool("oauth") || form.NewEmail == user.Email() {
		return (nil), (nil)
	}
	tokenCookie, err := c.Cookie(middleware.AuthCookieName)
	if err != nil {
		fmt.Println("Error getting cookie:", err)
		return (nil), (err)
	}

	data := structs.Map(form)
	headers := map[string]string{"Authorization": tokenCookie.Value}
	r := lib.PocketBaseRequest{
		Route:   "/users/request-email-change",
		Method:  "POST",
		Data:    data,
		Headers: headers,
	}
	resp, err := lib.NewPocketBaseRequest(r)
	fmt.Println(resp)

	return (resp), (err)
}
