package chain

import (
	"context"
	"fmt"
	"sync"
	"webhook/internal/entities"
)

type PubSub struct {
	topics map[string]chan <- *entities.Request
	mux sync.Mutex
}

func New() *PubSub {
	return &PubSub{
		topics: make(map[string]chan <- *entities.Request),
		mux:    sync.Mutex{},
	}
}

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

func (p *PubSub) Subscribe(ctx context.Context, topic string, messages chan<- *entities.Request) error {
	p.mux.Lock()
	defer p.mux.Unlock()
	
	p.topics[topic] = messages
	return nil
}
