package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sixserver/pkg/types"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(config types.Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error al conectar con Redis: %v", err)
	}
	log.Printf("Conectado a Redis en %s", config.Redis.Addr)
}
