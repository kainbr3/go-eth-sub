package storage

import (
	"context"
	"time"
)

type IStorage interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}
