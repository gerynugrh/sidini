package main

import (
	"github.com/gerynugrh/ubiquitous-octo-waffle/internal/forum"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"log"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		forum.RegisterRoute(e.Router)
		return nil
	})

	app.Dao().FindUserById()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
