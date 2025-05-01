package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), 
		Password: "",                      
		DB:       0,                      
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
