package injectors

import (
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/internal"
	"go.uber.org/zap"
)

func ProvideLogger(cfg *config.Config) *zap.SugaredLogger {
	appLogger := internal.NewLogger(cfg)
	appLogger.Logger = appLogger.Init()
	return appLogger.Logger
}
