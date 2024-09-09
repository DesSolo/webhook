package server

import (
	"net/http"
	"webhook/internal/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server HTTP server
type Server struct {
	mux chi.Router
}

// NewServer creates new HTTP server
func NewServer() *Server {
	return &Server{
		mux: chi.NewRouter(),
	}
}

// LoadRoutes loads routes
func (s *Server) LoadRoutes(p handlers.Publisher, sb handlers.Subscriber) {
	s.mux.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
	)

	s.mux.Route("/api/v1", func(r chi.Router) {

	})

	s.mux.Route("/webhook/{token}", func(r chi.Router) {
		r.Handle("/*", handlers.HandleWebhook(p))
	})

	s.mux.Handle("/ws", handlers.HandleWS(sb))

	s.mux.Get("/health", handlers.HandleHealth())
}

// Run starts HTTP server
func (s *Server) Run(address string) error {
	return http.ListenAndServe(address, s.mux)
}
