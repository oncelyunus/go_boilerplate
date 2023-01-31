package service

import (
	"context"

	"github.com/oncelyunus/go_boilerplate/pkg/app/repository/user"
	"go.uber.org/zap"
)

type UserService interface {
	SignUp(context context.Context, data *interface{}) (*interface{}, error)
}

type userService struct {
	logger         *zap.SugaredLogger
	userRepository user.UserRepository
}

func NewUserService(logger *zap.SugaredLogger, userRepository user.UserRepository) UserService {
	return &userService{
		logger:         logger,
		userRepository: userRepository,
	}
}

// SignUp implements UserService
func (us *userService) SignUp(context context.Context, data *interface{}) (*interface{}, error) {
	panic("unimplemented")
}
