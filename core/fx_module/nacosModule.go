package fx_module

import (
	"example.com/t/api/service"
	"go.uber.org/fx"
)

var FXNacosModule = fx.Module("fx-nacos-module",
	fx.Provide(service.NewNacosSerivce),
	fx.Invoke(func(s *service.NacosSerivce) {
		err := s.LoadAndWatchConfig()
		if err != nil {
			panic("NacosModule-Invoke: " + err.Error())
		}
	}),
	fx.Invoke(func(s *service.NacosSerivce) {
		err := s.RegisterService()
		if err != nil {
			panic("FXNacosModule: " + err.Error())
			return
		}
	}),
)
