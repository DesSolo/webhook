package simple

import (
	"net/http"
	"time"
	"webhook/internal/entities"
	"webhook/internal/responser"
)

// NewSimple new responser with simple params
func NewSimple(statusCode int, contentType, content string, timeout time.Duration) responser.ResponserFunc {
	return func(w http.ResponseWriter, r *entities.Request) error {
		if timeout > 0 {
			time.Sleep(timeout)
		}
		
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(statusCode)
		w.Write([]byte(content))
		return nil
	}
}
