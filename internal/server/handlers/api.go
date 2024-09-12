package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"webhook/internal/responser"
	"webhook/internal/responser/simple"
	"webhook/internal/service"

	"github.com/google/uuid"
)

// createChannelRequest create channel request
type createChannelRequest struct {
	Kind   string `json:"kind"`
	Simple struct {
		StatusCode  int    `json:"status_code"`
		ContentType string `json:"content_type"`
		Content     string `json:"content"`
		Timeout     int    `json:"timeout"`
	} `json:"simple"`
}

// Validate validate request
func (c *createChannelRequest) Validate() error {
	if c.Kind == "" {
		return fmt.Errorf("%w: missing kind", ErrInvalidRequest)
	}

	if c.Kind == "simple" {
		if c.Simple.StatusCode < 100 || c.Simple.StatusCode > 999 {
			return fmt.Errorf("%w: status code should be between 100 and 999", ErrInvalidRequest)
		}

		if c.Simple.Timeout > 30 {
			return fmt.Errorf("%w: timeout should be less than 30", ErrInvalidRequest)
		}
	}

	return nil
}

// createChannelResponse create response
type createChannelResponse struct {
	Token string `json:"token"`
}

// parseResponser parse responser from create request
func parseResponser(r *createChannelRequest) (responser.Responser, error) {
	// TODO: auto discovery
	switch r.Kind {
	case "simple":
		return simple.NewSimple(
			r.Simple.StatusCode,
			r.Simple.ContentType,
			r.Simple.Content,
			time.Duration(r.Simple.Timeout)*time.Second,
		), nil
	default:
		return nil, fmt.Errorf("kind %s not supported", r.Kind)
	}
}

// HandleChannelCreate handle create new channel
func HandleChannelCreate(ws *service.Webhook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createChannelRequest

		if err := bindJson(r, &req); err != nil {
			slog.ErrorContext(r.Context(),
				"fault bind request",
				"err", err,
			)

			respondJson(w, http.StatusBadRequest, errorMessage{
				Message: err.Error(),
			})

			return
		}
		defer r.Body.Close()

		responser, err := parseResponser(&req)
		if err != nil {
			slog.ErrorContext(r.Context(),
				"fault parse responser",
				"err", err,
			)

			respondJson(w, http.StatusBadRequest, errorMessage{
				Message: err.Error(),
			})

			return
		}

		token := uuid.New().String()

		ws.Register(token, responser)

		respondJson(w, http.StatusOK, createChannelResponse{
			Token: token,
		})
	}
}
