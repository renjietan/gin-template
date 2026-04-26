package fx_module

import (
	"example.com/t/core/sql_driver"
	"go.uber.org/fx"
)

var FXSqlliteModule = fx.Module("fx-sqlite-module",
	fx.Provide(sql_driver.NewSqliteDriver),
)
