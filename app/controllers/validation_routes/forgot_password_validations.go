package validation_routes

import (
	"encoding/json"
	"pb-starter/app/components/inputs"
	"pb-starter/app/forms"
	"pb-starter/app/lib"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterForgotPasswordValidationRoutes(e *core.ServeEvent, group echo.Group) {
	group.POST("/confirm-password-reset/password", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/confirm-password-reset/password", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/confirm-password-reset/password"}.Comp())
	})

	group.POST("/confirm-password-reset/password-confirm", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/confirm-password-reset/password-confirm", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/confirm-password-reset/password-confirm"}.Comp())
	})
}
