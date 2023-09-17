package service

import (
	"context"
	"testing"
	"time"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestVehicle_StartSimple(t *testing.T) {

	var err error

	// arrange
	cfg := sampleConfig()
	ctx := context.Background()
	deadline, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	v := &Vehicle{
		cfg: cfg,
	}
	// act
	go func() {
		err = v.Start(deadline)
		cancel()
	}()

	// assert
	assert.NoError(t, err, "o serviço devia iniciar sem problemas")
}

func TestVehicle_Stop(t *testing.T) {
	type fields struct {
		cfg *config.Cfg
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{{
		"no-error",
		fields{cfg: config.New()},
		args{ctx: context.Background()},
		false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vehicle{
				cfg: tt.fields.cfg,
			}
			if err := v.Stop(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Vehicle.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEmpty(t *testing.T) {
	var v Vehicle
	vehicle := New(&config.Cfg{})
	assert.NotNil(t, vehicle, "não deve estar vazia")
	assert.IsType(t, vehicle, &v, "deve ser um veículo")
}

func sampleConfig() *config.Cfg {
	cfg := &config.Cfg{
		Logger: zap.Must(zap.NewDevelopment()),
		Server: config.ServerSettings{
			HTTP: config.HTTPServerSettings{
				Endpoint: "localhost:10000",
			},
		},
	}
	return cfg
}
