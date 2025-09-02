package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

// Logger is a middleware function that logs incoming HTTP requests and their details
// using the provided zerolog.Logger instance. It captures information such as request
// time, remote address, request path, protocol, method, user agent, response status,
// bytes in/out, and elapsed time. It also handles panics and logs them as errors
func Logger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			startTime := time.Now()

			defer func() {
				endTime := time.Now()                 // End time
				elapsedTime := endTime.Sub(startTime) // Request time

				if r := recover(); r != nil && r != http.ErrAbortHandler {
					logger.Error().Interface("recover", r).Bytes("stack", debug.Stack()).Msg("incoming_request_panic")
					ww.WriteHeader(http.StatusInternalServerError)
				}

				logger.Info().Fields(map[string]interface{}{
					"time":         startTime.Format(time.RFC3339), // Format using RFC3339
					"remote_addr":  r.RemoteAddr,
					"path":         r.URL.Path,
					"proto":        r.Proto,
					"method":       r.Method,
					"user_agent":   r.UserAgent(),
					"status":       http.StatusText(ww.Status()),
					"status_code":  ww.Status(),
					"bytes_in":     r.ContentLength,
					"bytes_out":    ww.BytesWritten(),
					"elapsed_time": elapsedTime.String(),
				}).Msg("incoming_request")
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
