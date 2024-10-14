package controllers

import (
	"pb-starter/app/lib"
	"pb-starter/app/views/dashboard"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterDashboardRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/dashboard", func(c echo.Context) error {
		return lib.Render(c, dashboard.Dashboard(c))
	})
}
