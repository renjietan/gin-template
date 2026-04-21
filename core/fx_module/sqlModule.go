package fx_module

import (
	"example.com/t/api/service"
	"example.com/t/core"
	"go.uber.org/fx"
)

var FXSqlModule = fx.Module("fx-sql-module",
	// 初始化数据库
	fx.Provide(core.NewGormConfig),
	fx.Provide(core.NewMysql),
	// 自动同步 数据库 表
	fx.Provide(service.NewMigrationService),
	fx.Invoke(func(migrationService *service.MigrationService) {
		migrationService.StartMigrate()
	}),
)
