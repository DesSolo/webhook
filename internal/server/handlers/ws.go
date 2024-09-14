package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"webhook/internal/entities"

	"github.com/gorilla/websocket"
)

// Subscriber interface for pubsub request
type Subscriber interface {
	Subscribe(ctx context.Context, token string, messages chan<- *entities.Request) error
}

// HandleWS handle websocket
func HandleWS(s Subscriber) http.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		channelID := r.URL.Query().Get("channel")
		if channelID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger := slog.With(
			"channel", channelID,
		)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.ErrorContext(r.Context(), 
				"failed to upgrade websocket", 
				"err", err,
			)

			return
		}
		defer conn.Close()

		logger.InfoContext(r.Context(), "connected")

		ch := make(chan *entities.Request)

		go func() {
			if err := s.Subscribe(r.Context(), channelID, ch); err != nil {
				logger.ErrorContext(r.Context(), 
					"failed to subscribe", 
					"err", err,
				)

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()

		for {
			select {
			case <-r.Context().Done():
				logger.InfoContext(r.Context(), "disconnected")
				return
			case req := <-ch:
				if err := conn.WriteJSON(req); err != nil {
					logger.ErrorContext(r.Context(), 
						"failed to write message", 
						"err", err,
					)
					
					return
				}
			}
		}
	}
}
