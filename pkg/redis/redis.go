package pkgredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(dsn string) *RedisClient {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	return &RedisClient{
		client: redis.NewClient(opts),
	}
}

func (obj *RedisClient) Set(key string, value any, duration time.Duration) bool {
	res := obj.client.Set(context.Background(), key, value, duration)
	if res.Err() != nil {
		return false
	}
	return true
}

func (obj *RedisClient) Get(key string) any {
	res := obj.client.Get(context.Background(), key)
	if res.Err() != nil {
		return nil
	}
	return res.Val()
}
