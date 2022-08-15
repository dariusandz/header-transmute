package reader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingParser(t *testing.T) {
	// given
	const key1 = "key1"
	const value1 = "value1"
	const key2 = "key2"
	const value2 = "value2"
	const key3 = "key3"
	const value3 = "value3"
	mapping := fmt.Sprintf("%s:%s\n%s:%s\n%s:%s", key1, value1, key2, value2, key3, value3)

	// when
	parsedMapping, _ := ParseHeaderMapping(mapping)

	// then
	assert.Equal(t, parsedMapping[key1], value1)
	assert.Equal(t, parsedMapping[key2], value2)
	assert.Equal(t, parsedMapping[key3], value3)
}

func TestMappingParserEmpty(t *testing.T) {
	// when
	parsedMapping, _ := ParseHeaderMapping("")

	// expect
	assert.Equal(t, map[string]string{}, parsedMapping)
}

func TestMappingParserInvalid(t *testing.T) {
	// given
	mapping := fmt.Sprintf("%s|%s", "key1", "value1")

	// when
	_, err := ParseHeaderMapping(mapping)

	// then
	assert.Error(t, err)
}
