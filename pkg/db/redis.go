package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func ConnectRedis(url string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: url,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis")
	}
	log.Println("Connected to Redis")
}
