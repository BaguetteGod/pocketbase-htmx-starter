package validation_routes

import (
	"encoding/json"
	"pb-starter/app/components/inputs"
	"pb-starter/app/forms"
	"pb-starter/app/lib"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterRegisterValidationRoutes(e *core.ServeEvent, group echo.Group) {
	group.POST("/register/email", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/register/email", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/register/email"}.Comp())
	})

	group.POST("/register/username", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/register/username", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/register/username"}.Comp())
	})

	group.POST("/register/display-name", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "displayname", Label: "Display name", Value: form.DisplayName, HxPost: "/register/display-name", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "displayname", Label: "Display name", Value: form.DisplayName, HxPost: "/register/display-name"}.Comp())
	})

	group.POST("/register/password", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/register/password", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/register/password"}.Comp())
	})

	group.POST("/register/password-confirm", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/register/password-confirm", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/register/password-confirm"}.Comp())
	})
}
