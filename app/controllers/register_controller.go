package controllers

import (
	"encoding/json"
	"pb-starter/app/components/toast"
	"pb-starter/app/controllers/validation_routes"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	"pb-starter/app/views/register"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterRegisterRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/register", func(c echo.Context) error {
		return lib.Render(c, register.RegisterPage(forms.RegisterFormValue{}, ""))
	})

	group.POST("/register", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, register.RegisterPage(form, string(formErrors)))
		}

		err = lib.Register(e, c, form.Email, form.Username, form.DisplayName, form.Password)
		if err != nil {
			return err
		}

		// data := map[string]interface{}{"email": form.Email}
		data := structs.Map(form)
		r := lib.PocketBaseRequest{
			Route:  "/users/request-verification",
			Method: "POST",
			Data:   data,
		}
		resp, _ := lib.NewPocketBaseRequest(r)
		if resp.StatusCode != 204 {
			toast.New(c, toast.Toast{
				Level:   "DANGER",
				Title:   "Something went wrong",
				Message: "We weren't able to send a verification email. Please try again.",
			})
			return lib.Render(c, register.RegisterSuccessPage(form))

		}

		return lib.Render(c, register.RegisterSuccessPage(form))
	})

	validation_routes.RegisterRegisterValidationRoutes(e, group)
}
