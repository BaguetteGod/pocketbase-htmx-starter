package main

import (
	"date-rate/app/pocketbase/extensions"
	"date-rate/app/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/dist", "app/internal/assets/dist")

		routes.Init(e, app)
		extensions.SetPocketbaseExtensions(e, app)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
