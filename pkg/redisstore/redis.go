package redisstore

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx context.Context
)

func InitRedis() {
	Ctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		Rdb = nil
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
