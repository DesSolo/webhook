package pubsub

import (
	"context"
	"webhook/internal/entities"
)

type PubSub interface {
	Publish(ctx context.Context, topic string, r *entities.Request) error
	Subscribe(ctx context.Context, topic string, messages chan<- *entities.Request) error
	// TODO: add close
}
