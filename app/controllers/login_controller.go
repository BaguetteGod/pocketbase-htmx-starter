package controllers

import (
	"encoding/json"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	"pb-starter/app/views/login"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLoginRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/login")
	})

	group.GET("/login", func(c echo.Context) error {
		return lib.Render(c, login.LoginPage(forms.LoginFormValue{}, ""))
	})

	group.POST("/login", func(c echo.Context) error {
		form := forms.GetLoginFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, login.LoginPage(form, string(formErrors)))
		}

		user, _ := e.App.Dao().FindAuthRecordByEmail("users", form.Email)
		if !user.Verified() {
			errorMap := map[string]interface{}{
				"verified": "Please verify your email",
			}
			formErrors, _ := json.Marshal(errorMap)
			return lib.Render(c, login.LoginPage(form, string(formErrors)))
		}

		err = lib.Login(e, c, form.Email, form.Password)
		if err != nil {
			return err
		}

		return c.Redirect(302, "/dashboard")
	})

	group.POST("/login/oauth2", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		lib.SetCookie(c, json_map["token"].(string))

		return c.Redirect(302, "/login")
	})
}
