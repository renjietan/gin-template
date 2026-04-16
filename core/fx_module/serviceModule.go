package fx_module

import (
	"example.com/t/api/controller"
	"example.com/t/api/service"
	"go.uber.org/fx"
)

var ApiModule = fx.Module("ApiModule",
	// 公用模块
	fx.Provide(service.NewUploadService),
	// 模块 - config
	fx.Provide(service.NewConfigService),
	fx.Provide(controller.NewConfigController),
	fx.Invoke(func(cfgController *controller.ConfigController) {
		cfgController.RegisterRouter()
	}),
)
