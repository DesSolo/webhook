package handlers

import (
	"encoding/json"
	"net/http"
)

// respondJson responds JSON data
func respondJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
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

// bindJson unmarshal json
func bindJson(r *http.Request, v Validator) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}