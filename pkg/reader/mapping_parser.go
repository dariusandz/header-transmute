package reader

import (
	"bufio"
	"errors"
	"regexp"
	"strings"
)

/*
ParseHeaderMapping reads string of format
	mapFromValue1:mapToValue1
	mapFromValue2:mapToValue2
and returns a map of key value pairs split by first ':'
*/
func ParseHeaderMapping(headerMapping string) (map[string]string, error) {
	headerRegex := regexp.MustCompile("^.*:.*$")
	m := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(headerMapping))
	for scanner.Scan() {
		line := scanner.Text()
		if matches := headerRegex.MatchString(line); !matches {
			return nil, errors.New("header mapping must be in form of key value pairs split by ':'")
		} else {
			mapping := parseMapping(line)
			m[mapping.ValueFrom] = mapping.ValueTo
		}
	}

	return m, nil
}

func parseMapping(line string) mapping {
	headerMapping := strings.SplitN(line, ":", 2)
	return mapping{
		ValueFrom: strings.TrimSpace(headerMapping[0]),
		ValueTo:   strings.TrimSpace(headerMapping[1]),
	}
}

type mapping struct {
	ValueFrom string
	ValueTo   string
}
