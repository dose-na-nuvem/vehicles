package service

import (
	"context"
	"fmt"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/dose-na-nuvem/vehicles/pkg/server"
)

type Vehicle struct {
	cfg  *config.Cfg
	http *server.HTTP
}

func New(cfg *config.Cfg) *Vehicle {
	return &Vehicle{
		cfg:  cfg,
		http: &server.HTTP{},
	}
}

func (v *Vehicle) Start(ctx context.Context) error {
	var err error

	v.cfg.Logger.Info("inicializando...")
	v.cfg.Logger.Info("🚗💨 vehicles running")

	vh := server.NewVehiclesHandler(v.cfg.Logger)
	v.http, err = server.NewHTTP(v.cfg, vh)
	if err != nil {
		return fmt.Errorf("falha ao configurar o servidor HTTP: %w", err)
	}
	err = v.http.Start(ctx)
	if err != nil {
		return fmt.Errorf("falha ao iniciar o servidor HTTP: %w", err)
	}

	return nil
}

func (v *Vehicle) Stop(ctx context.Context) error {
	v.cfg.Logger.Info("terminando...")
	return nil
}
