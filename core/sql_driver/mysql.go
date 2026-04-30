package sql_driver

import (
	"time"

	"example.com/t/types"
	"example.com/t/utility"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDriver(config *gorm.Config, appConfig *types.AppConfig) (*gorm.DB, error) {
	template := `{{Username}}:{{Password}}@tcp({{Host}}:{{Port}})/{{Database}}?charset=utf8mb4&parseTime=True&loc=Local`
	dns := utility.StringByTemplate(template, appConfig.Mysql)
	db, err := gorm.Open(mysql.Open(dns), config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(32)
	sqlDB.SetMaxOpenConns(512)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
