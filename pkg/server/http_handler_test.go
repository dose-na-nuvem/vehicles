package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestVehicleHandlerError(t *testing.T) {
	// prepare
	core, logs := observer.New(zap.InfoLevel)
	logger := zap.New(core)
	h := NewVehiclesHandler(logger)
	writer := httptest.NewRecorder()

	// test
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, "", nil)
	require.NoError(t, err)
	h.ServeHTTP(writer, req)

	// verify
	response := writer.Result()
	defer response.Body.Close()
	assert.Equal(t, response.StatusCode, http.StatusNotImplemented)

	assert.Len(t, logs.All(), 1)
	assert.Contains(t, logs.All()[0].Message, infoStarting)

}
