package memory

import (
	"context"
	"sync"
	"webhook/internal/responser"
	"webhook/internal/storage"
)

// Storage memory storage
type Storage struct {
	data map[string]responser.Responser
	mux  *sync.RWMutex
}

// New constructor
func New() *Storage {
	return &Storage{
		data: make(map[string]responser.Responser),
		mux:  &sync.RWMutex{},
	}
}

// SaveResponser save responser	to memory
func (s *Storage) SaveResponser(_ context.Context, token string, rs responser.Responser) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[token] = rs
	return nil
}

// LoadResponser load responser from memory
func (s *Storage) LoadResponser(_ context.Context, token string) (responser.Responser, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	rs, ok := s.data[token]
	if !ok {
		return nil, storage.ErrNotExist
	}

	return rs, nil
}
