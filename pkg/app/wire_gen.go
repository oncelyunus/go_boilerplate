// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/injectors"
)

// Injectors from wire.go:

func NewApp(cfg *config.Config) (*application, error) {
	sugaredLogger := injectors.ProvideLogger(cfg)
	mux := injectors.ProvideRouter(cfg)
	server := injectors.ProvideHttpServer(cfg, mux)
	healthCheckHandler := injectors.ProvideHealthCheck(sugaredLogger)
	mongodbConnector, err := injectors.ProvideMongoDB(cfg, sugaredLogger)
	if err != nil {
		return nil, err
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(cfg, mongodbConnector)
	userRepository := injectors.ProvideUserRepository(baseMongoRepo)
	userService := injectors.ProvideUserService(sugaredLogger, userRepository)
	authHandler := injectors.ProvideAuthHandler(sugaredLogger, userService)
	appApplication := &application{
		config:             cfg,
		logger:             sugaredLogger,
		router:             mux,
		httpServer:         server,
		healthCheckHandler: healthCheckHandler,
		authHandler:        authHandler,
	}
	return appApplication, nil
}
