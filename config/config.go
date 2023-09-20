package config

import "go.uber.org/zap"

type Cfg struct {
	Server ServerSettings `mapstructure:"server"`

	Logger *zap.Logger
}

type GRPCServerSettings struct {
	Endpoint string `mapstructure:"endpoint"`
}

type TLSSettings struct {
	CertFile    string `mapstructure:"cert_file"`
	CertKeyFile string `mapstructure:"cert_key_file"`
	Insecure    bool   `mapstructure:"insecure"`
}


type ServerSettings struct {
	GRPC GRPCServerSettings `mapstructure:"grpc"`
	TLS  *TLSSettings       `mapstructure:"tls"`
}

func New() *Cfg {
	cfg := &Cfg{
		Logger: zap.Must(zap.NewDevelopment()),
		Server: ServerSettings{
			TLS: &TLSSettings{},
		},
	}
	return cfg
}
