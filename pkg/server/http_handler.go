package server

import (
	"net/http"

	"go.uber.org/zap"
)

const (
	infoStarting = "tratando requisições"
)

var _ http.Handler = (*VehiclesHandler)(nil)

type VehiclesHandler struct {
	logger *zap.Logger
}

func NewVehiclesHandler(logger *zap.Logger) *VehiclesHandler {
	return &VehiclesHandler{
		logger: logger,
	}
}

func (vh *VehiclesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vh.logger.Info(infoStarting)
	switch r.Method {
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
