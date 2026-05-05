package fx_module

import (
	"example.com/t/core/component/sql_driver"
	"go.uber.org/fx"
)

var FxGormConfigModule = fx.Module("fx-gorm-config-module",
	fx.Provide(sql_driver.NewGormConfig),
)
