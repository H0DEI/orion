package handlers

import "net/http"

// Health responds with a simple OK to indicate the service is running.
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
