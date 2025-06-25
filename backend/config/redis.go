package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}
	Rdb = redis.NewClient(&redis.Options{Addr: redisURL})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Gagal Konek ke Redis", err)
	}

	log.Println("Redis Connected to", redisURL)
}
