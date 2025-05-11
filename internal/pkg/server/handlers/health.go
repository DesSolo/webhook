package handlers

import (
	"log/slog"
	"net/http"
)

// HandleHealth health check
func HandleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte("OK")); err != nil {
			slog.Error("failed to write message", "err", err)
		}
	}
}
