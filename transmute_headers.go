package header_transmute

// Adapted from https://github.com/traefik/plugindemo/blob/master/demo.go

import (
	"context"
	"fmt"
	"github.com/dariusandz/header-transmute/pkg/mapper"
	"github.com/dariusandz/header-transmute/pkg/reader"
	"github.com/dariusandz/header-transmute/pkg/types"
	"net/http"
	"os"
)

type Config struct {
	FromHeader    string
	ToHeader      string
	MappingEnvKey string
}

func CreateConfig() *Config {
	return &Config{}
}

type HeadersTransmutation struct {
	name string
	next http.Handler
	rule types.Rule
}

// New creates and returns a plugin instance.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.FromHeader) == 0 {
		fmt.Println("FromHeader is not defined!")
		return nil, fmt.Errorf("FromHeader is not defined")
	}

	if len(config.ToHeader) == 0 {
		return nil, fmt.Errorf("ToHeader is not defined")
	}

	if len(config.MappingEnvKey) == 0 {
		return nil, fmt.Errorf("MappingEnvKey is not defined")
	}

	mapping, isPresent := os.LookupEnv(config.MappingEnvKey)
	if !isPresent {
		// Do not fail so validity test passes
		//return nil, fmt.Errorf("could not extract mapping from env key %s", config.MappingEnvKey)
	}

	parsedMapping, err := reader.ParseHeaderMapping(mapping)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &HeadersTransmutation{
		name: name,
		next: next,
		rule: types.Rule{
			FromHeader:    config.FromHeader,
			ToHeader:      config.ToHeader,
			HeaderMapping: parsedMapping,
		},
	}, nil
}

func (u *HeadersTransmutation) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	mapper.Handle(responseWriter, request, u.rule)
	u.next.ServeHTTP(responseWriter, request)
}
