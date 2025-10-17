package app

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loadRoutes() http.Handler {

	// Initialize logger
	l := log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout,
	})

	// Initialize router using chi
	router := chi.NewRouter()

	// Middleware for logging requests
	router.Use(middleware.Logger(l))

	// Middleware for CORS
	router.Use(middleware.CorsMiddleware())

	// Add Prometheus /metrics endpoint
	router.Handle("/metrics", promhttp.Handler())

	return router
}