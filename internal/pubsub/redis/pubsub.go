package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"webhook/internal/entities"

	goredis "github.com/redis/go-redis/v9"
)

// PubSub based on redis
type PubSub struct {
	client goredis.UniversalClient
}

// New constructor
func New(client goredis.UniversalClient) *PubSub {
	return &PubSub{
		client: client,
	}
}

// Publish publish request to topic
func (p *PubSub) Publish(ctx context.Context, token string, r *entities.Request) error {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("fault marshal request: %w", err)
	}

	return p.client.Publish(ctx, token, data).Err()
}

// Subscribe subscribe to topic
func (p *PubSub) Subscribe(ctx context.Context, token string, messages chan<- *entities.Request) error {
	sub := p.client.PSubscribe(ctx, token)

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return fmt.Errorf("fault receive message: %w", err)
		}

		var r entities.Request
		if err := json.Unmarshal([]byte(msg.Payload), &r); err != nil {
			return fmt.Errorf("fault unmarshal message: %w", err)
		}

		messages <- &r
	}
}

// Close close pubsub
func (p *PubSub) Close() error {
	return p.client.Close()
}
