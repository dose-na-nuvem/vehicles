package config

import (
	"time"

	"go.uber.org/zap"
)

type Cfg struct {
	Logger *zap.Logger
	Server ServerSettings `mapstructure:"server"`
}

type ServerSettings struct {
	HTTP HTTPServerSettings `mapstructure:"http"`
}

type HTTPServerSettings struct {
	Endpoint          string        `mapstructure:"endpoint"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
}

func New() *Cfg {
	cfg := &Cfg{
		Logger: zap.Must(zap.NewDevelopment()),
		Server: ServerSettings{},
	}
	return cfg
}
