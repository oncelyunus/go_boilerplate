package injectors

import (
	"github.com/oncelyunus/go_boilerplate/pkg/app/repository/user"
	"github.com/oncelyunus/go_boilerplate/pkg/app/service"
	"go.uber.org/zap"
)

func ProvideUserService(logger *zap.SugaredLogger, userRepository user.UserRepository) service.UserService {
	return service.NewUserService(logger, userRepository)
}
