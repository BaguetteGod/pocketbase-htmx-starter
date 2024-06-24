package controllers

import (
	"date-rate/app/lib"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterLogoutRoutes(e *core.ServeEvent, group echo.Group) {
	group.POST("/logout", func(c echo.Context) error {
		lib.SetCookie(c, "")

		return c.Redirect(302, "/login")
	})
}
