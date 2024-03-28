package redis

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/redis/go-redis/v9"
	"gitlab.com/revolutionize1/foward-api/internal/app"
)

var Instance *redis.Client

func Init() {
	opts, err := createRedisOptions()
	if err != nil {
		log.Fatalf("Failed to create Redis options: %v", err)
	}

	Instance = redis.NewClient(opts)

	if err := testRedisConnection(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func createRedisOptions() (*redis.Options, error) {
	return &redis.Options{
		Addr:           fmt.Sprintf("redis:%s", app.Instance.Config.Redis.Port),
		Password:       app.Instance.Config.Redis.Password,
		DB:             0,
		MaxActiveConns: 50,
		MaxIdleConns:   25,
		PoolSize:       25,
	}, nil
}

func testRedisConnection() error {
	_, err := Instance.Ping(context.Background()).Result()
	return err
}
