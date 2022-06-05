package store

import (
	"log"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type redisDatabase struct {
	client *redis.Client
}

type redisCache struct {
	cache *cache.Cache
}

type KVStore interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

type RedisError struct {
	err string
}

func (e *RedisError) Error() string {
	return e.err
}

func logError(err error) (string, error) {
	log.Printf("Error: %v", err)
	return "[Error]", &RedisError{}
}
