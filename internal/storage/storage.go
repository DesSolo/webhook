package storage

import (
	"context"
	"webhook/internal/responser"
)

// ResponseStorage storage for responser
type ResponseStorage interface {
	SaveResponser(ctx context.Context, token string, rs responser.Responser) error
	LoadResponser(ctx context.Context, token string) (responser.Responser, error)
}
