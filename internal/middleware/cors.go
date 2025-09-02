package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CorsMiddleware is a middleware function that sets up CORS (Cross-Origin Resource Sharing)
// for the HTTP server. It allows requests from any origin and supports common HTTP methods
func CorsMiddleware() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
