package routes

import (
	"date-rate/app/controllers"
	middleware "date-rate/app/middelware"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Init(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.Use(middleware.LoadAuthContextFromCookie(app))

	unauthenticatedGroup := e.Router.Group("", middleware.UnauthGuard)
	controllers.RegisterLoginRoutes(e, *unauthenticatedGroup)
	controllers.RegisterRegisterRoutes(e, *unauthenticatedGroup)
	controllers.RegisterForgotPasswordRoutes(e, *unauthenticatedGroup)
	controllers.RegisterEmailVerifiedRoutes(e, *unauthenticatedGroup)

	authenticatedGroup := e.Router.Group("", middleware.AuthGuard)
	controllers.RegisterLogoutRoutes(e, *authenticatedGroup)
	controllers.RegisterDashboardRoutes(e, *authenticatedGroup)
}
