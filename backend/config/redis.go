package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Gagal Konek ke Redis", err)
	}

	log.Println("Redis Connected")
}
