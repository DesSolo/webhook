package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// respondJSON responds JSON data
func respondJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to write message",
			"err", err,
		)
	}
}

// errorMessage for respond error
type errorMessage struct {
	Message string `json:"error"`
}

// Validator interface
type Validator interface {
	// Validate validate request
	Validate() error
}

// bindJSON unmarshal json
func bindJSON(r *http.Request, v Validator) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}
