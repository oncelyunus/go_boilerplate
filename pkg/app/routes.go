package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/oncelyunus/go_boilerplate/pkg/internal"
)

func (app *application) routes() {
	router := app.router
	router.Use(internal.ChiLogger(app.logger))
	router.Get("/health", app.healthCheckHandler.HealthCheck)
	router.Mount("/auth", app.authRouter())
}

func (app *application) authRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/signup", app.authHandler.SignUp)

	return r
}
