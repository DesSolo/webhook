package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"webhook/internal/entities"
	"webhook/internal/responser"
	"webhook/internal/responser/simple"
)

// TODO: move to config
var defaultResponser = simple.New(http.StatusOK, "application/text", "", 0)

// Publisher interface for publish request in topic
type Publisher interface {
	Publish(ctx context.Context, topic string, r *entities.Request) error
}

// ResponseStorage interface for store responser
type ResponseStorage interface {
	SaveResponser(ctx context.Context, token string, rs responser.Responser) error
	LoadResponser(ctx context.Context, token string) (responser.Responser, error)
}

// Webhook webhook service
type Webhook struct {
	publisher Publisher
	storage   ResponseStorage
}

// NewWebhook constructor
func NewWebhook(p Publisher, s ResponseStorage) *Webhook {
	return &Webhook{
		publisher: p,
		storage:   s,
	}
}

// Register register new responser for token
func (w *Webhook) Register(ctx context.Context, token string, rs responser.Responser) error {
	if err := w.storage.SaveResponser(ctx, token, rs); err != nil {
		return fmt.Errorf("fault save responser: %w", err)
	}

	return nil
}

// Responser get responser by token
func (w *Webhook) Responser(ctx context.Context, token string) responser.Responser {
	rs, err := w.storage.LoadResponser(ctx, token)
	if err != nil {
		slog.Error("fault load responser", "err", err)
		return defaultResponser
	}

	return rs
}

// Handle handle webhook request
func (w *Webhook) Handle(ctx context.Context, rw http.ResponseWriter, req *entities.Request) error {
	rs := w.Responser(ctx, req.Token)

	if err := rs.Response(rw, req); err != nil {
		return fmt.Errorf("fault response request: %w", err)
	}

	if err := w.publisher.Publish(ctx, req.Token, req); err != nil {
		return fmt.Errorf("fault publish request: %w", err)
	}

	return nil
}
