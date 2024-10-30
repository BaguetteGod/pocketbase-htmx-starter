package validation_routes

import (
	"encoding/json"
	"pb-starter/app/components/inputs"
	"pb-starter/app/forms"
	"pb-starter/app/lib"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterProfileValidationRoutes(e *core.ServeEvent, group echo.Group) {
	group.POST("/profile/email", func(c echo.Context) error {
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{
				Name:   "newEmail",
				Label:  "Email address",
				Value:  form.NewEmail,
				HxPost: "/profile/email",
				Error:  string(formErrors),
			}.Comp())
		}

		return lib.Render(c, inputs.Text{
			Name:   "newEmail",
			Label:  "Email address",
			Value:  form.NewEmail,
			HxPost: "/profile/email",
		}.Comp())
	})

	group.POST("/profile/username", func(c echo.Context) error {
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{
				Name:   "username",
				Value:  form.Username,
				HxPost: "/profile/username",
				Error:  string(formErrors),
			}.Comp())
		}

		return lib.Render(c, inputs.Text{
			Name:   "username",
			Value:  form.Username,
			HxPost: "/profile/username",
		}.Comp())
	})

	group.POST("/profile/display-name", func(c echo.Context) error {
		form := forms.GetProfileFormValue(c)
		err := form.Validate(e, c)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{
				Name:   "displayname",
				Label:  "Display name",
				Value:  form.DisplayName,
				HxPost: "/profile/display-name",
				Error:  string(formErrors),
			}.Comp())
		}

		return lib.Render(c, inputs.Text{
			Name:   "displayname",
			Label:  "Display name",
			Value:  form.DisplayName,
			HxPost: "/profile/display-name",
		}.Comp())
	})
}
