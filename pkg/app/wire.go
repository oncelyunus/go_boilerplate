//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/injectors"
)

func NewApp(cfg *config.Config) (*application, error) {
	panic(
		wire.Build(
			injectors.ProvideLogger,
			injectors.ProvideRouter,
			injectors.ProvideHttpServer,
			injectors.ProvideHealthCheck,

			//repository
			injectors.ProvideBaseMongoRepo,
			injectors.ProvideMongoDB,
			injectors.ProvideUserRepository,

			//service
			injectors.ProvideUserService,

			//handler
			injectors.ProvideAuthHandler,
			wire.Struct(new(application), "*"),
		),
	)
}
