package mapper

import (
	"github.com/dariusandz/header-transmute/pkg/types"
	"net/http"
)

func Handle(_ http.ResponseWriter, req *http.Request, rule types.Rule) {
	for headerName, headerValues := range req.Header {
		if headerMatches := rule.FromHeader == headerName; !headerMatches {
			continue
		}

		req.Header.Del(headerName)

		for _, headerValue := range headerValues {
			if _, mappingExists := rule.HeaderMapping[headerValue]; mappingExists {
				req.Header.Add(rule.ToHeader, rule.HeaderMapping[headerValue])
				continue
			}

			req.Header.Add(rule.ToHeader, headerValue)
		}
	}
}
