package util

import (
	"time"

	"github.com/go-redis/redis"
)

type redisCapsule struct {
	cache *redis.Client
}

type Cacher interface {
	Set(key string, pair string, value interface{}, duration time.Duration)
	Delete(key string, pair string)
	Get(key string, pair string) string
}

func RedisUtility(cache *redis.Client) Cacher {
	return &redisCapsule{
		cache: cache,
	}
}

func (c *redisCapsule) Set(k string, p string, v interface{}, d time.Duration) {
	c.cache.Del(k + ":" + p)
	if d != 0*time.Second {
		c.cache.SetNX(k+":"+p, v, d)
	} else {
		c.cache.Set(k+":"+p, v, 0)
	}
}

func (c *redisCapsule) Delete(k string, p string) {
	c.cache.Del(k + ":" + p)
}

func (c *redisCapsule) Get(k string, p string) string {
	val, err := c.cache.Get(k + ":" + p).Result()
	if err != nil {
		return ""
	}
	return val
}
