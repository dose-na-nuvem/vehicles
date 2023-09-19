package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/dose-na-nuvem/vehicles/config"
	"go.uber.org/zap"
)

const (
	httpStarting = "servidor http iniciando"
	httpStopping = "servidor http finalizando"
)

type HTTP struct {
	logger *zap.Logger

	srv *http.Server
}

func NewHTTP(cfg *config.Cfg, vehiclesHandler http.Handler) (*HTTP, error) {

	mux := http.NewServeMux()
	mux.Handle("/", vehiclesHandler)

	srv := &http.Server{
		Addr:    cfg.Server.HTTP.Endpoint,
		Handler: mux,
	}

	return &HTTP{
		logger: cfg.Logger,
		srv:    srv,
	}, nil
}

func (h *HTTP) Start(ctx context.Context) error {
	h.logger.Info(httpStarting, zap.String("endpoint", h.srv.Addr))
	err := h.srv.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		h.logger.Error("falha ao iniciar o servidor HTTP", zap.Error(err))
	}

	return nil
}

func (h *HTTP) Shutdown(ctx context.Context) error {
	h.logger.Info(httpStopping)
	// We received an interrupt signal, shut down.
	if err := h.srv.Shutdown(ctx); err != nil {
		// Erro ao fechar ouvintes (listeners) ou tempo limite (timeout) de contexto.
		return err
	}

	return nil
}
