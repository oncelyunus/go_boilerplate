package injectors

import (
	"github.com/oncelyunus/go_boilerplate/pkg/app/controller"
	"github.com/oncelyunus/go_boilerplate/pkg/app/service"
	"go.uber.org/zap"
)

func ProvideAuthHandler(logger *zap.SugaredLogger, userService service.UserService) *controller.AuthHandler {
	return controller.NewAuthHandler(logger, userService)
}
