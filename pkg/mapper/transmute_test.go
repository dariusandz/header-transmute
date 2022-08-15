package mapper

import (
	"github.com/stretchr/testify/assert"
	"header-transmute/pkg/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

const headerToTransmuteFrom = "Some-Header"
const headerToTransmuteTo = "Some-Other-Header"
const oldHeaderValue = "oldValue"
const newHeaderValue = "newValue"

func TestTransmuteHandler(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost", nil)
	req.Header.Set(headerToTransmuteFrom, oldHeaderValue)

	rule := types.Rule{
		FromHeader:    headerToTransmuteFrom,
		ToHeader:      headerToTransmuteTo,
		HeaderMapping: map[string]string{oldHeaderValue: newHeaderValue},
	}

	// when
	Handle(recorder, req, rule)

	// then
	assert.Equal(t, "", req.Header.Get(headerToTransmuteFrom))
	assert.Equal(t, newHeaderValue, req.Header.Get(headerToTransmuteTo))
}

func TestTransmuteHandlerWithNoMapping(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost", nil)
	req.Header.Set(headerToTransmuteFrom, oldHeaderValue)

	rule := types.Rule{
		FromHeader:    headerToTransmuteFrom,
		ToHeader:      headerToTransmuteTo,
		HeaderMapping: map[string]string{},
	}

	// when
	Handle(recorder, req, rule)

	// then
	assert.Equal(t, "", req.Header.Get(headerToTransmuteTo))
	assert.Equal(t, oldHeaderValue, req.Header.Get(headerToTransmuteFrom))
}