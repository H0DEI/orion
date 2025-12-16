package main

import (
	"log"
	"net/http"
	"time"

	"github.com/H0DEI/orion/internal/config"
	"github.com/H0DEI/orion/internal/httpapi"
)

func main() {
	// Load application configuration from environment variables.
	cfg := config.Load()

	// Build HTTP router.
	router := httpapi.NewRouter()

	// Configure HTTP server with sensible timeouts.
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("orion api starting on http://localhost%s", cfg.HTTPAddr)

	// Start HTTP server.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
