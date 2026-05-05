package sql_driver

import (
	"example.com/t/types"
	"gorm.io/gorm"
)

type SqlManager struct {
	MySqlDriver  *gorm.DB
	SqliteDriver *gorm.DB
	AppConfig    *types.AppConfig
}
