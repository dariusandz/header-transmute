package mapper

import (
	"header-transmute/pkg/types"
	"net/http"
)

func Handle(_ http.ResponseWriter, req *http.Request, rule types.Rule) {
	for headerName, headerValues := range req.Header {
		if headerMatches := rule.FromHeader == headerName; !headerMatches {
			continue
		}

		for _, headerValue := range headerValues {
			if _, mappingExists := rule.HeaderMapping[headerValue]; !mappingExists {
				continue
			}

			req.Header.Del(headerName)
			req.Header.Set(rule.ToHeader, rule.HeaderMapping[headerValue])
		}
	}
}
