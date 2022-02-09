package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, velue *entity.Post) {

}

func (cache *redisCache) Get(key string) *entity.Post {

}
