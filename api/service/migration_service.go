package service

import (
	"example.com/t/api/model"
	"example.com/t/types"
	"gorm.io/gorm"
)

type MigrationService struct {
	db        *gorm.DB
	appConfig *types.AppConfig
}

func NewMigrationService(db *gorm.DB, appConfig *types.AppConfig) *MigrationService {
	return &MigrationService{
		db:        db,
		appConfig: appConfig,
	}
}

func (s *MigrationService) StartMigrate() {
	go func() {
		s.TableMigration()
	}()
}

func (s *MigrationService) TableMigration() {
	s.db.AutoMigrate(&model.Config{})
	s.db.AutoMigrate(&model.LoginLog{})
}
