package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string, pass string) *RedisCache {
	c := &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
		}),
	}

	return c
}

func (c *RedisCache) Set(key, value interface{}, ttl int64) error {
	err := c.client.Set(key.(string), value, time.Duration(ttl)*time.Second).Err()

	return err
}

func (c *RedisCache) Get(key interface{}) (interface{}, error) {
	val, err := c.client.Get(key.(string)).Result()

	return val, err
}
