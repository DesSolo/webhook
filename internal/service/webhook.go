package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"webhook/internal/entities"
	"webhook/internal/pubsub"
	"webhook/internal/responser"
	"webhook/internal/responser/simple"
)

// TODO: move to config
var defaultResponser = simple.NewSimple(http.StatusOK, "application/text", "", 0)


// Webhook webhook service
type Webhook struct {
	pubsub pubsub.PubSub
	// TODO: move to response storage
	responses map[string]responser.Responser
	mux       sync.Mutex
}

// NewWebhook constructor
func NewWebhook(ps pubsub.PubSub) *Webhook {
	return &Webhook{
		pubsub:    ps,
		responses: make(map[string]responser.Responser),
		mux:       sync.Mutex{},
	}
}

// Register register new responser for token
func (w *Webhook) Register(token string, rs responser.Responser) {
	w.mux.Lock()
	defer w.mux.Unlock()

	w.responses[token] = rs
}

// Responser get responser by token
func (w *Webhook) Responser(token string) responser.Responser {
	w.mux.Lock()
	defer w.mux.Unlock()

	rs, ok := w.responses[token]
	if !ok {
		return defaultResponser
	}

	return rs
}

// Handle handle webhook request
func (w *Webhook) Handle(ctx context.Context, rw http.ResponseWriter, req *entities.Request) error {
	rs := w.Responser(req.Token)

	if err := rs.Response(rw, req); err != nil {
		return fmt.Errorf("fault response request: %w", err)
	}

	if err := w.pubsub.Publish(ctx, req.Token, req); err != nil {
		return fmt.Errorf("fault publish request: %w", err)
	}

	return nil
}
