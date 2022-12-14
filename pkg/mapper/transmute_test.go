package mapper_test

import (
	"github.com/dariusandz/header-transmute/pkg/mapper"
	"github.com/dariusandz/header-transmute/pkg/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	headerToTransmuteFrom = "Some-Header"
	headerToTransmuteTo   = "Some-Other-Header"
	oldHeaderValue        = "oldValue"
	newHeaderValue        = "newValue"
)

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
	mapper.Handle(recorder, req, rule)

	// then
	assert.Equal(t, 0, len(req.Header.Values(headerToTransmuteFrom)))
	assert.Equal(t, newHeaderValue, req.Header.Get(headerToTransmuteTo))
}

func TestTransmuteHandlerWithMultipleValues(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost", nil)
	const retainedValue = "retainedValue"
	req.Header.Add(headerToTransmuteFrom, retainedValue)
	req.Header.Add(headerToTransmuteFrom, oldHeaderValue)

	rule := types.Rule{
		FromHeader:    headerToTransmuteFrom,
		ToHeader:      headerToTransmuteTo,
		HeaderMapping: map[string]string{oldHeaderValue: newHeaderValue},
	}

	// when
	mapper.Handle(recorder, req, rule)

	// then
	assert.Equal(t, 0, len(req.Header.Values(headerToTransmuteFrom)))
	assert.Equal(t, 2, len(req.Header.Values(headerToTransmuteTo)))
	assert.Contains(t, req.Header.Values(headerToTransmuteTo), retainedValue)
	assert.Contains(t, req.Header.Values(headerToTransmuteTo), newHeaderValue)
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
	mapper.Handle(recorder, req, rule)

	// then
	assert.Equal(t, 0, len(req.Header.Values(headerToTransmuteFrom)))
	assert.Equal(t, 1, len(req.Header.Values(headerToTransmuteTo)))
	assert.Equal(t, oldHeaderValue, req.Header.Get(headerToTransmuteTo))
}
