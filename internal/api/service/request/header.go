package request

import (
	"net/http"
)

func parseHeader(header header) http.Header {
	headerMap := make(http.Header)

	for k, v := range header {
		key := http.CanonicalHeaderKey(k)
		headerMap[key] = append(headerMap[key], v)
	}
	return headerMap
}
