package redis

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedisDB() error {
	Redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if Rdb == nil {
		return errors.New("redis connect Failed")
	}
	Rdb = Redisdb
	return nil
}
