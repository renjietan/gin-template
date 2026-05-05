package fx_module

import (
	"example.com/t/core/component/swagger"
	"go.uber.org/fx"
)

var FXSwaggerModule = fx.Module("fx-swagger-module",
	fx.Provide(swagger.NewSwaggerManager),
	fx.Invoke(func(sw *swagger.SwaggerManager) {
		sw.InitRouter()
	}),
)
