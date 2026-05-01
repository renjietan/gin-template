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

func NewSqlManager(app_config *types.AppConfig) *SqlManager {
	//if app_config.Sqlite.Enable == true {
	//
	//}
	return nil
}
