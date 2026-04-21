package fx_module

import (
	"example.com/t/api/controller"
	"example.com/t/api/service"
	"example.com/t/core"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var ApiModule = fx.Module("fx-api-module",
	// 公用模块
	fx.Invoke(func(appserver *core.AppServer) { // swagger 路由
		appserver.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// 根据环境变量来判断是否需要 禁用 swagger
		appserver.Engine.GET("/swagger1/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "APP_DEBUG"))
	}),
	fx.Provide(service.NewUploadService), // 图片上传
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
