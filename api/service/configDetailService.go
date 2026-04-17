package service

import (
	"sync"

	"example.com/t/api/entity"
	"gorm.io/gorm"
)

type ConfigDetailService struct {
	db   *gorm.DB
	lock sync.Mutex
}

func NewConfigDetailService(db *gorm.DB) *ConfigDetailService {
	return &ConfigDetailService{db: db, lock: sync.Mutex{}}
}

func (c *ConfigDetailService) getById(id string) []entity.ConfigDetailEntity {
	var details []entity.ConfigDetailEntity
	c.db.Joins("Config").Find(&details)
	return details
}
