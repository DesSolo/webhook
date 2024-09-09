package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondJson responds with JSON
func respondJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
