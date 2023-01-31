package controller

import (
	"net/http"

	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/internal"
	"go.uber.org/zap"
)

type HealthCheckHandler struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func NewHealthCheckHandler(logger *zap.SugaredLogger) *HealthCheckHandler {
	return &HealthCheckHandler{logger: logger}
}

func (h *HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("I am here HealthCheckHandler")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	internal.ToJSON(&internal.GenericResponse{Status: true, Message: ""}, w)
}
