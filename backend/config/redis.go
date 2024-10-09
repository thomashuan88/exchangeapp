package config

import (
	"exchangeapp/backend/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to initialize redis, error: %v", err)
	}

	log.Println("Redis initialized")

	global.RedisDB = RedisClient
}
