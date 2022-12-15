package redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	cacheSizeLimit int = 10000
)

type RedisCache[T any] struct {
	cache                 *cache.Cache
	defaultExpirationTime time.Duration
}

func NewRedisCache[T any](client *redis.Client, defaultExpirationTime time.Duration) *RedisCache[T] {
	return &RedisCache[T]{
		cache: cache.New(&cache.Options{
			Redis:      client,
			LocalCache: cache.NewTinyLFU(cacheSizeLimit, defaultExpirationTime),
		}),
		defaultExpirationTime: defaultExpirationTime,
	}
}

func (rc *RedisCache[T]) Set(ctx context.Context, key string, v T) error {
	return rc.setWithTime(ctx, rc.defaultExpirationTime, key, v)
}

func (rc *RedisCache[T]) setWithTime(ctx context.Context, duration time.Duration, key string, v T) error {
	return rc.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: v,
		TTL:   duration,
	})
}

func (rc *RedisCache[T]) Get(ctx context.Context, key string) (*T, error) {
	var value T
	if err := rc.cache.Get(ctx, key, &value); err != nil {
		return nil, err
	}

	return &value, nil
}

func (rc *RedisCache[T]) Delete(ctx context.Context, key string) error {
	return rc.cache.Delete(ctx, key)
}
