package controller

import (
	"net/http"

	"github.com/oncelyunus/go_boilerplate/pkg/app/service"
	"go.uber.org/zap"
)

type AuthHandler struct {
	logger      *zap.SugaredLogger
	UserService service.UserService
}

func NewAuthHandler(logger *zap.SugaredLogger, userService service.UserService) *AuthHandler {
	return &AuthHandler{logger: logger, UserService: userService}
}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ah.logger.Info("I am here SignUp")
	var i interface{}
	ah.UserService.SignUp(r.Context(), &i)
}
