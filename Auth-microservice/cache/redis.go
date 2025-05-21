package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // use default URL
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// check connection to redis
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err 
	}

	log.Print("Connected to redis")
	return nil
}