package fx_module

import (
	"example.com/t/core/logger"
	"example.com/t/core/nacos"
	"go.uber.org/fx"
)

var FXNacosModule = fx.Module("fx-nacos-module",
	fx.Provide(nacos.NewNacosSerivce),
	fx.Invoke(func(s *nacos.NacosSerivce) {
		err := s.LoadAndWatchConfig()
		if err != nil {
			e := logger.BeautifyJsonStr(err)
			logger.L().Error(e)
		}
	}),
	fx.Invoke(func(s *nacos.NacosSerivce) {
		err := s.RegisterService()
		if err != nil {
			e := logger.BeautifyJsonStr(err)
			logger.L().Error(e)
		}
	}),
)
