package main

import (
	"github.com/ThatsMrTalbot/example"
	"github.com/ThatsMrTalbot/example/handlers"
)

// Application is the main enrty point for the program
type Application struct {
	*handlers.Index `inject:""`
	Config          *example.Config `inject:""`
}

// Start application
func (app *Application) Start() error {
	app.Config.Servers.Start(app)
	return nil
}
