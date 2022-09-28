package log

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-logr/logr"
)

func Middleware(logger logr.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				t2 := time.Now()

				// Recover and record stack traces in case of a panic.
				if rec := recover(); rec != nil {
					logger.Error(
						fmt.Errorf("internal server error"), "server error occurred", "recoverInfo", rec, "debugStack",
						debug.Stack(),
					)
					http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}

				// Log request.
				logger.Info(
					"incoming request", "requestID", middleware.GetReqID(r.Context()), "remoteIP",
					r.RemoteAddr, "url",
					r.URL.Path, "proto", r.Proto, "method",
					r.Method, "userAgent", r.Header.Get("User-Agent"), "status", ww.Status(), "latencyMs",
					float64(t2.Sub(t1).Nanoseconds())/1000000.0, "bytesIn", r.Header.Get("Content-Length"), "bytesOut",
					ww.BytesWritten(),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
