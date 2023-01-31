package injectors

import (
	"fmt"
	"net/http"

	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/go-chi/chi"
)

func ProvideHttpServer(cfg *config.Config, router *chi.Mux) *http.Server {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	return httpServer
}
