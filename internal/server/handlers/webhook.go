package handlers

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"
	"webhook/internal/entities"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func newFromRequest(r *http.Request, token string, data []byte) *entities.Request {
	// TODO: duration, size
	return &entities.Request{
		UUID:    uuid.New().String(),
		Token:   token,
		Date:    time.Now().Format(time.RFC3339),
		IP:      r.RemoteAddr,
		Method:  r.Method,
		URI:     r.RequestURI,
		Query:   r.URL.Query().Encode(),
		Headers: r.Header,
		Data:    data,
	}
}

// Publisher interface for pubsub request
type Publisher interface {
	Publish(ctx context.Context, token string, req *entities.Request) error
}

// HandleWebhook handle webhook
func HandleWebhook(p Publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")
		if token == "" {
			slog.DebugContext(r.Context(),
				"missing token",
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			slog.DebugContext(r.Context(),
				"fault read request body",
				"err", err,
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		req := newFromRequest(r, token, data)

		if err := p.Publish(r.Context(), token, req); err != nil {
			slog.DebugContext(r.Context(),
				"fault publish request",
				"err", err,
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJson(w, http.StatusOK, req)
	}
}
