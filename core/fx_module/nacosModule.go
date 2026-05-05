package fx_module

import (
	logger2 "example.com/t/core/component/logger"
	"example.com/t/core/component/nacos"
	"go.uber.org/fx"
)

var FXNacosModule = fx.Module("fx-nacos-module",
	fx.Provide(nacos.NewNacosSerivce),
	fx.Invoke(func(s *nacos.NacosSerivce) {
		err := s.LoadAndWatchConfig()
		if err != nil {
			e := logger2.BeautifyJsonStr(err)
			logger2.L().Error(e)
		}
	}),
	fx.Invoke(func(s *nacos.NacosSerivce) {
		err := s.RegisterService()
		if err != nil {
			e := logger2.BeautifyJsonStr(err)
			logger2.L().Error(e)
		}
	}),
)
