package controllers

import (
	"pb-starter/app/lib"
	"pb-starter/app/views/email_verified"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterEmailVerifiedRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/confirm-verification/:token", func(c echo.Context) error {
		return lib.Render(c, email_verified.EmailVerifiedSuccessPage())
	})
}
