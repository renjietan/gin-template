package redis_driver

import (
	"context"
	"fmt"
	"time"

	"example.com/t/types"
	"github.com/go-redis/redis/v8"
)

func NewRedisDriver(app_config *types.AppConfig) (*redis.Client, error) {
	url := fmt.Sprintf("%s:%d", app_config.Redis.Host, app_config.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:        url,
		Password:    app_config.Redis.Password,
		DB:          app_config.Redis.DB,
		PoolSize:    app_config.Redis.PoolSize,
		PoolTimeout: time.Duration(app_config.Redis.PoolTimeout) * time.Second,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	app_config.Redis.Enable = true
	return client, nil
}
