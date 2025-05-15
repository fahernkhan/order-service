package resource

import (
	// golang package
	"context"
	"fmt"
	"log"
	"order-service/config"

	// external package
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis init redis by given cfg pointer of config.Config.
//
// It returns pointer of redis.Client when successful.
// Otherwise, nil pointer of redis.Client will be returned.
func InitRedis(cfg *config.Config) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connect to redis: %v", err)
	}

	log.Println("Connected to Redis")
	return RedisClient
}
