package fx_module

import (
	"example.com/t/api/service"
	"example.com/t/core/logger"
	"go.uber.org/fx"
)

var FXNacosModule = fx.Module("fx-nacos-module",
	fx.Provide(service.NewNacosSerivce),
	fx.Invoke(func(s *service.NacosSerivce) {
		err := s.LoadAndWatchConfig()
		if err != nil {
			e := logger.BeautifyJsonStr(err)
			logger.L().Error(e)
		}
	}),
	fx.Invoke(func(s *service.NacosSerivce) {
		err := s.RegisterService()
		if err != nil {
			e := logger.BeautifyJsonStr(err)
			logger.L().Error(e)
		}
	}),
)
