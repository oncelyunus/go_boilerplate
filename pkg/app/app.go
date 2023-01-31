package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/app/controller"
	"go.uber.org/zap"
)

type Application interface {
	Init() error
	StartServer() error
}

type application struct {
	config     *config.Config
	logger     *zap.SugaredLogger
	router     *chi.Mux
	httpServer *http.Server

	healthCheckHandler *controller.HealthCheckHandler
	authHandler        *controller.AuthHandler
}

func (app *application) Init() error {
	app.routes()
	return nil
}

func (app *application) StartServer() error {
	go func() {
		err := app.httpServer.ListenAndServe()
		if err != nil {
			app.logger.Errorf("Failed to start http server at port :%d \n", app.config.Server.Port)
		}
	}()
	app.logger.Infof("Server started at port with graceful shutdown :%d \n", app.config.Server.Port)
	app.startGracefulShutdownChan()
	return nil
}

func (app *application) startGracefulShutdownChan() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT)
	// interrupt signal sent from terminal
	signal.Notify(gracefulStop, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(gracefulStop, syscall.SIGTERM)
	app.logger.Info("Enabled Graceful Shutdown \n")

	select {
	case sig := <-gracefulStop:
		app.logger.Warnf("caught sig: %+v \n", sig)
	}

	app.shutDownProcesses()
}

func (app *application) shutDownProcesses() {

	// We received an interrupt signal, server shut down.
	if err := app.httpServer.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		app.logger.Error("HTTP server Shutdown \n")
	}
	app.logger.Infof("Http server shut down finished. \n")
	// this sleep is to give buffer so that any pending process can completes gracefully,.
	app.logger.Infof("Wait for %v to finish processing \n", 3*time.Second)
	time.Sleep(3 * time.Second)
	app.logger.Infof("Shutting down. \n")
	os.Exit(0)
}
