//go:build wireinject
// +build wireinject

package app

import (
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/injectors"
	"github.com/google/wire"
)

func NewApp(cfg *config.Config) (*application, error) {
	panic(
		wire.Build(
			injectors.ProvideLogger,
			injectors.ProvideRouter,
			injectors.ProvideHttpServer,
			injectors.ProvideHealthCheck,
			wire.Struct(new(application), "*"),
		),
	)
}
