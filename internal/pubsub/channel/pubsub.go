package channel

import (
	"context"
	"fmt"
	"sync"
	"webhook/internal/entities"
)

// PubSub based on channels
type PubSub struct {
	topics map[string]chan <- *entities.Request
	mux sync.Mutex
}

// New constructor
func New() *PubSub {
	return &PubSub{
		topics: make(map[string]chan <- *entities.Request),
		mux:    sync.Mutex{},
	}
}

// Publish publish request to topic
func (p *PubSub) Publish(ctx context.Context, topic string, r *entities.Request) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	ch, ok := p.topics[topic]
	if !ok {
		return fmt.Errorf("unknown topic: %s", topic)
	}

	ch <- r
	return nil
}

// Subscribe subscribe to topic
func (p *PubSub) Subscribe(ctx context.Context, topic string, messages chan<- *entities.Request) error {
	p.mux.Lock()
	defer p.mux.Unlock()
	
	p.topics[topic] = messages
	return nil
}
