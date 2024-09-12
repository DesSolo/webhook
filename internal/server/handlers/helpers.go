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

type errorMessage struct {
	Message string `json:"error"`
}

type Validator interface {
	Validate() error
}

func bindJson(r *http.Request, v Validator) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}