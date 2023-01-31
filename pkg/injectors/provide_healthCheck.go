package injectors

import (
	"github.com/oncelyunus/go_boilerplate/pkg/app/controller"
	"go.uber.org/zap"
)

func ProvideHealthCheck(logger *zap.SugaredLogger) *controller.HealthCheckHandler {
	configHandler := controller.NewHealthCheckHandler(logger)
	return configHandler
}
