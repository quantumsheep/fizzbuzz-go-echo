package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client

	ctx context.Context
}

func NewRedisCache(address string) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisCache) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}
