package controllers

import (
	"date-rate/app/lib"
	"date-rate/app/views/dashboard"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterDashboardRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/dashboard", func(c echo.Context) error {
		info := apis.RequestInfo(c)
		record := info.AuthRecord
		fmt.Println(record.Get("avatar"))
		return lib.Render(c, dashboard.Dashboard(c))
	})
}
