package service

import (
	"sync"

	"gorm.io/gorm"
)

type ConfigService struct {
	db   *gorm.DB
	lock sync.Mutex
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{db: db, lock: sync.Mutex{}}
}
