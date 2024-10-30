package controllers

import (
	"encoding/json"
	"pb-starter/app/components/toast"
	"pb-starter/app/controllers/validation_routes"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	"pb-starter/app/views/forgot_password"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterForgotPasswordRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/forgot-password", func(c echo.Context) error {
		return lib.Render(c, forgot_password.ForgotPasswordPage(forms.ForgotPasswordFormValue{}, ""))
	})

	group.POST("/forgot-password", func(c echo.Context) error {
		form := forms.GetForgotPasswordFormValue(c)
		err := form.Validate()
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, forgot_password.ForgotPasswordPage(forms.ForgotPasswordFormValue{}, string(formErrors)))
		}

		data := structs.Map(form)
		r := lib.PocketBaseRequest{
			Route:  "/users/request-password-reset",
			Method: "POST",
			Data:   data,
		}
		resp, _ := lib.NewPocketBaseRequest(r)
		if resp.StatusCode != 204 {
			toast.New(c, toast.Toast{
				Level:   toast.DANGER,
				Title:   "Something went wrong",
				Message: "We weren't able to send a password reset email. Please try again.",
			})
			return lib.Render(c, forgot_password.ForgotPasswordPage(form, ""))
		}

		return lib.Render(c, forgot_password.ForgotPasswordSuccessPage())
	})

	group.GET("/confirm-password-reset/:token", func(c echo.Context) error {
		token := c.PathParam("token")
		user, _ := e.App.Dao().FindAuthRecordByToken(token, e.App.Settings().RecordPasswordResetToken.Secret)
		if user == nil {
			return c.Redirect(302, "/login")
		}
		form := forms.ConfirmPasswordResetFormValue{Token: token}

		return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, ""))
	})

	group.POST("/confirm-password-reset/:token", func(c echo.Context) error {
		form := forms.GetConfirmPasswordResetForm(c)
		err := form.Validate()
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, string(formErrors)))
		}

		data := structs.Map(form)
		r := lib.PocketBaseRequest{
			Route:  "/users/confirm-password-reset",
			Method: "POST",
			Data:   data,
		}
		resp, _ := lib.NewPocketBaseRequest(r)
		if resp.StatusCode != 204 {
			toast.New(c, toast.Toast{Level: toast.DANGER, Title: "Link is invalid or has expired"})
			return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, ""))
		}

		return lib.Render(c, forgot_password.ConfirmPasswordSuccessPage())
	})

	validation_routes.RegisterForgotPasswordValidationRoutes(e, group)
}
