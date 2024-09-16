package simple

import (
	"encoding/json"
	"net/http"
	"time"
	"webhook/internal/entities"
	"webhook/internal/responser"
)

func init() {
	responser.Register(&Simple{})
}

// Simple simple responser
type Simple struct {
	// Response status code
	StatusCode int
	// Response content type
	ContentType string
	// Response content
	Content string
	// Response timeout
	Timeout time.Duration
}

// New constructor
func New(statusCode int, contentType string, content string, timeout time.Duration) *Simple {
	return &Simple{
		StatusCode:  statusCode,
		ContentType: contentType,
		Content:     content,
		Timeout:     timeout,
	}
}

// Response apply response for request
func (s *Simple) Response(w http.ResponseWriter, r *entities.Request) error {
	if s.Timeout > 0 {
		time.Sleep(s.Timeout)
	}

	w.Header().Set("Content-Type", s.ContentType)
	w.WriteHeader(s.StatusCode)
	w.Write([]byte(s.Content))
	return nil
}

// Kind responser kind
func (s *Simple) Kind() string {
	return "simple"
}

// MarshalBinary marshal
func (s *Simple) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary unmarshal
func (s *Simple) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
