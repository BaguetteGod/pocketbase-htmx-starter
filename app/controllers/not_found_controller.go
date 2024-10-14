package controllers

import (
	"pb-starter/app/lib"
	notfound "pb-starter/app/views/not_found"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterNotFounRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("", func(c echo.Context) error {
		return lib.Render(c, notfound.NotFound())
	})
}
