package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(redis *redis.Client) *Redis {
	return &Redis{
		client: redis,
	}
}

func (r *Redis) Save(ctx context.Context, key string, data any) error {
	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *Redis) Get(ctx context.Context, key string, data any) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrCacheNil
		}
		return err
	}
	err = json.Unmarshal([]byte(val), data)
	if err != nil {
		return err
	}
	return nil
}
