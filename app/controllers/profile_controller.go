package controllers

import (
	"date-rate/app/forms"
	"date-rate/app/lib"
	"date-rate/app/views/profile"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterProfileRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/profile", func(c echo.Context) error {
		user := apis.RequestInfo(c).AuthRecord
		form := forms.ProfileFormValue{
			Username: user.Username(),
			Email:    user.Email(),
		}
		return lib.Render(c, profile.Profile(c, form, ""))
	})
}
