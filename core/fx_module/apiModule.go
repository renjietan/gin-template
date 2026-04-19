package fx_module

import (
	"example.com/t/api/controller"
	"example.com/t/api/service"
	"example.com/t/core"
	"go.uber.org/fx"
)

var ApiModule = fx.Module("fx-api-module",
	// 公用模块
	fx.Provide(service.NewUploadService),
	// 模块 - config
	fx.Provide(service.NewConfigService),
	fx.Provide(controller.NewConfigController),
	fx.Invoke(func(cfgController *controller.ConfigController, core *core.AppServer) {
		cfgController.RegisterConfigRouters()
	}),
	// 模块 - config detail
	fx.Provide(service.NewConfigDetailService),
	fx.Provide(controller.NewConfigDetailController),
	fx.Invoke(func(cfgDetailController *controller.ConfigDetailsController) {
		cfgDetailController.RegisterConfigDetailRouters()
	}),
)
