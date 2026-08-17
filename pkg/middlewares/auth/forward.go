package auth

import (
	"net/http"
)

// copyHeaders copies headers from the auth response to the request.
// It ensures that we do not modify shared state and that headers are scoped to the request.
func copyHeaders(dst, src http.Header, keys []string) {
	for _, key := range keys {
		if values := src.Values(key); len(values) > 0 {
			// Use a fresh slice to avoid potential reference sharing issues
			newValues := make([]string, len(values))
			copy(newValues, values)
			dst[http.CanonicalHeaderKey(key)] = newValues
		}
	}
}
