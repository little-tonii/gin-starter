package config

import (
	"context"
	"fmt"
	"gin-starter/internal/shared/constant"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient     *redis.Client
	redisClientOnce sync.Once
)

func InitializeRedisCahing() error {
	var err error
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%d",
				constant.Environment.REDIS_CACHING_HOST,
				constant.Environment.REDIS_CACHING_PORT,
			),
			Password: constant.Environment.REDIS_CACHING_PASSWORD,
			DB:       constant.Environment.REDIS_CACHING_DB,
		})
		ping := redisClient.Ping(context.Background())
		if ping.Err() != nil {
			err = ping.Err()
		}
	})
	return err
}

func GetRedisClient() *redis.Client {
	return redisClient
}
