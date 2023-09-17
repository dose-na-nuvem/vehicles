package service

import (
	"context"

	"github.com/dose-na-nuvem/vehicles/config"
)

type Vehicle struct {
	cfg *config.Cfg
}

func New(cfg *config.Cfg) *Vehicle {
	return &Vehicle{
		cfg: cfg,
	}
}

func (v *Vehicle) Start(ctx context.Context) error {
	v.cfg.Logger.Info("inicializando...")
	v.cfg.Logger.Info("🚗💨 vehicles running")
	return nil
}

func (v *Vehicle) Stop(ctx context.Context) error {
	v.cfg.Logger.Info("terminando...")
	return nil
}
