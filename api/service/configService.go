package service

import (
	"fmt"
	"sync"

	dto "example.com/t/api/DTO"
	"example.com/t/api/entity"
	"gorm.io/gorm"
)

type ConfigService struct {
	db   *gorm.DB
	lock sync.Mutex
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{db: db, lock: sync.Mutex{}}
}

func (c *ConfigService) Insert(d dto.ConfigDTO) entity.ConfigEntity {
	c.lock.Lock()
	defer c.lock.Unlock()
	var res entity.ConfigEntity
	c.db.Create(&entity.ConfigEntity{
		Name:  d.Name,
		Value: d.Value,
	}).Scan(&res)
	return res
}

func (c *ConfigService) InsertMany(d dto.ConfigsDTO) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	var res []entity.ConfigEntity
	for _, v := range d.Items {
		temp := entity.ConfigEntity{
			Name:  v.Name,
			Value: v.Value,
		}
		res = append(res, temp)
	}
	tx := c.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %w", tx.Error)
	}
	if err := tx.CreateInBatches(res, 100); err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	return tx.Commit().Error
}

func (c *ConfigService) upload(d dto.ConfigDTO) entity.ConfigEntity {
	c.lock.Lock()
	defer c.lock.Unlock()
	var res entity.ConfigEntity
	c.db.Create(&entity.ConfigEntity{
		Name:  d.Name,
		Value: d.Value,
	}).Scan(&res)
	return res
}
