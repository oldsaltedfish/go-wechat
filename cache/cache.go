package cache

import (
	"context"
	"errors"
)

type Cache interface {
	Save(ctx context.Context, key string, data any) error
	Get(ctx context.Context, key string, data any) error
}

var ErrCacheNil = errors.New("cache is nil")
