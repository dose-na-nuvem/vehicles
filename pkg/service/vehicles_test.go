package service

import (
	"context"
	"testing"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/stretchr/testify/assert"
)

func TestVehicle_Start(t *testing.T) {
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
	}{
		{
			"no-error",
			fields{cfg: config.New()},
			args{context.Background()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vehicle{
				cfg: tt.fields.cfg,
			}
			if err := v.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Vehicle.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
	assert.NotNil(t, vehicle, "even so, should be not empty")
	assert.IsType(t, vehicle, &v, "deve ser um veículo")
}
