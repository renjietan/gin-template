package fx_module

import (
	"example.com/t/api/service"
	"example.com/t/core/component/sql_driver"
	"go.uber.org/fx"
)

var FXMySqlModule = fx.Module("fx-sql-module",
	// 初始化数据库
	fx.Provide(sql_driver.NewMysqlDriver),
	// 自动同步 数据库 表
	fx.Provide(service.NewMigrationService),
	fx.Invoke(func(migrationService *service.MigrationService) {
		migrationService.StartMigrate()
	}),
)
