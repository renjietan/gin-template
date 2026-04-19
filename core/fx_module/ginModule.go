package fx_module

import (
	"example.com/t/core"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var GinModule = func(debug bool) fx.Option {
	return fx.Module("fx-gin-module",
		fx.Provide(core.NewAppServer),
		// 注意: 注册路由前，必须保证 core.AppServer 中注册了 中间件
		ApiModule,
		fx.Invoke(func(appserver *core.AppServer, db *gorm.DB, log *logrus.Logger) {
			appserver.Run(debug, log)
		}),
	)
}
