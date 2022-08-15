package header_transmute_test

import (
	"context"
	"fmt"
	plugin "github.com/dariusandz/header-transmute"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const headerToTransmuteFrom = "Some-Header"
const headerToTransmuteTo = "Some-Other-Header"
const newHeaderValue = "newValue"
const envKey = "Environment-Variable-Key"

func TestHeadersTransmute(t *testing.T) {
	// given
	mapping := fmt.Sprintf("oldValue:%s", newHeaderValue)

	os.Setenv(envKey, mapping)
	defer os.Unsetenv(envKey)

	cfg := plugin.Config{
		FromHeader:    headerToTransmuteFrom,
		ToHeader:      headerToTransmuteTo,
		MappingEnvKey: envKey,
	}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, _ := plugin.New(ctx, next, &cfg, "plugin-headers")
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	req.Header.Add(headerToTransmuteFrom, "oldValue")

	// when
	handler.ServeHTTP(recorder, req)

	// then
	assert.Equal(t, "", req.Header.Get(headerToTransmuteFrom))
	assert.Equal(t, newHeaderValue, req.Header.Get(headerToTransmuteTo))
}
