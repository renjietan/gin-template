package redis_driver

import (
	"context"
	"fmt"
	"time"

	"example.com/t/types"
	"github.com/go-redis/redis/v8"
)

func NewRedisDriver(app_config *types.AppConfig) (*redis.Client, error) {
	url := fmt.Sprintf("%s:%d", app_config.Redis_Host, app_config.Redis_Port)
	client := redis.NewClient(&redis.Options{
		Addr:        url,
		Password:    app_config.RedisConfig.Redis_Password,
		DB:          app_config.RedisConfig.Redis_DB,
		PoolSize:    app_config.RedisConfig.Redis_PoolSize,
		PoolTimeout: time.Duration(app_config.RedisConfig.Redis_PoolTimeout) * time.Second,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
