package httpapi

import (
	"net/http"

	"github.com/H0DEI/orion/internal/httpapi/handlers"
)

// NewRouter builds the HTTP routes for the API.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	return mux
}
