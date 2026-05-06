package sql_driver

import (
	"example.com/t/types"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DriverManager struct {
	MySqlDriver  *gorm.DB
	SqliteDriver *gorm.DB
	RedisDriver  *redis.Client
	AppConfig    *types.AppConfig
}

func NewDriverManager() (*DriverManager, error) {
	return &DriverManager{}, nil
}
