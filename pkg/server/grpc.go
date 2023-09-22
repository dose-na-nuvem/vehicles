package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/dose-na-nuvem/vehicles/config"
	//"github.com/dose-na-nuvem/vehicles/proto/vehicle"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var errNoTLSConfig = errors.New("servidor sem configuração de TLS")

type GRPC struct {
	//customer.UnimplementedCustomerServer
	logger   *zap.Logger
	grpc     *grpc.Server
	listener net.Listener
}

func NewGRPC(cfg *config.Cfg ) (*GRPC, error) {
	lis, err := net.Listen("tcp", cfg.Server.GRPC.Endpoint)
	if err != nil {
		return nil, err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	grpc := &GRPC{
		logger:   cfg.Logger,
		//store:    store,
		grpc:     grpcServer,
		listener: lis,
	}

	//customer.RegisterCustomerServer(grpcServer, grpc)

	return grpc, nil
}

func buildServerOptions(cfg *config.Cfg) ([]grpc.ServerOption, error) {
	var opts []grpc.ServerOption

	// tls certificates
	if cfg.Server.TLS.CertFile != "" && cfg.Server.TLS.CertKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(cfg.Server.TLS.CertFile,
			cfg.Server.TLS.CertKeyFile)
		if err != nil {
			return nil, fmt.Errorf("%s, %w", errNoTLSConfig, err)
		}

		opts = append(opts, grpc.Creds(creds))
	} else {
		if cfg.Server.TLS.Insecure {
			cfg.Logger.Info("Servidor sem configurações de TLS! Este servidor está inseguro!")
		} else {
			return nil, errNoTLSConfig
		}
	}

	return opts, nil
}

func (g *GRPC) Start(_ context.Context, chErr chan error) {
	g.logger.Info("iniciando servidor gRPC")
	go func() {
		err := g.grpc.Serve(g.listener)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			chErr <- fmt.Errorf("falha ao iniciar o servidor GRPC: %w", err)
		}
	}()
}

func (g *GRPC) Shutdown(_ context.Context) error {
	g.logger.Info("finalizando servidor gRPC")
	g.grpc.GracefulStop()
	return nil
}
