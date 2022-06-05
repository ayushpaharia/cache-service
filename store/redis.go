package store

import (
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var redis_cache redisCache

func CreateKVStore() (KVStore, error) {
	redis_db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := redis_db.Ping(redis_db.Context()).Result()
	if err != nil {
		log.Fatal(err)
		return nil, &RedisError{}
	}

	redis_cache = redisCache{
		cache.New(&cache.Options{
			Redis: redis_db,
		}),
	}

	return &redisDatabase{client: redis_db}, nil
}

func (r *redisDatabase) Set(key string, value string) (string, error) {
	if err := redis_cache.cache.Set(&cache.Item{
		Ctx:   r.client.Context(),
		Key:   "ayush" + key,
		Value: value,
		TTL:   time.Second * 3000,
	}); err != nil {
		return logError(err)
	}

	return key, nil
}

func (r *redisDatabase) Get(key string) (string, error) {
	var value string
	if err := redis_cache.cache.Get(r.client.Context(), key, &value); err != nil {
		return logError(err)
	}

	return value, nil
}
