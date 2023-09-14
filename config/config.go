package config

import "go.uber.org/zap"

type Cfg struct {
	Logger *zap.Logger
}

func New() *Cfg {
	cfg := &Cfg{
		Logger: zap.Must(zap.NewDevelopment()),
	}
	return cfg
}
