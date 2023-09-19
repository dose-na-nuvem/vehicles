package server

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTP_BlockingStartSuccessfulAndShutdown(t *testing.T) {
	var err error
	// prepare
	ctx := context.Background()
	deadline, cancel := context.WithTimeout(ctx, time.Microsecond*100)

	listener, port, err := GetListenerWithFallback(3, 43678)
	require.NoError(t, err, "não foi possivel alocar uma porta livre")
	listener.Close()
	freePortEndpoint := fmt.Sprintf("localhost:%d", port)

	cfg := config.New()
	cfg.Server.HTTP.Endpoint = freePortEndpoint

	srv := &http.Server{
		Addr: freePortEndpoint,
	}

	h := &HTTP{
		logger: cfg.Logger,
		srv:    srv,
	}

	// act
	go func() {
		err = h.Start(deadline)
		cancel()
	}()

	// assert
	assert.NoError(t, err, "falha ao iniciar o servidor http")

	// assert
	err = h.Shutdown(ctx)
	assert.NoError(t, err, "não deve ter erro se foi inicializado corretamente")
}
