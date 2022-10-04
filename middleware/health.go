package middleware

import (
	"net/http"
	"strings"
)

// Heartbeat is a middleware to return a 200 status for a health check of the server.
func Heartbeat(endpoint string) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Hermitboard is healthy! ^-^"))
				return
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
