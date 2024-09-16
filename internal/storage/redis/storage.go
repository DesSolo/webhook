package redis

import (
	"context"
	"fmt"
	"webhook/internal/responser"

	"github.com/redis/go-redis/v9"
)

// Storage redis storage
type Storage struct {
	client redis.UniversalClient
}

// New constructor
func New(client redis.UniversalClient) *Storage {
	return &Storage{
		client: client,
	}
}

// SaveResponser save responser to redis
func (s *Storage) SaveResponser(ctx context.Context, token string, rs responser.Responser) error {
	return s.client.Set(ctx, token, newMetadata(rs), 0).Err()
}

// LoadResponser load responser from redis
func (s *Storage) LoadResponser(ctx context.Context, token string) (responser.Responser, error) {
	var m metadata
	if err := s.client.Get(ctx, token).Scan(&m); err != nil {
		return nil, fmt.Errorf("fault load responser: %w", err)
	}

	return m.Responser, nil
}
