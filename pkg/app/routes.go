package app

import "github.com/oncelyunus/go_boilerplate/pkg/internal"

func (app *application) routes() {
	router := app.router
	router.Use(internal.ChiLogger(app.logger))
	router.Get("/health", app.healthCheckHandler.HealthCheck)
}
