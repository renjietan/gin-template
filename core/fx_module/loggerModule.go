package fx_module

import (
	"example.com/t/core/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var LoggerModule = fx.Module("logger",
	fx.Provide(logger.NewLogger),
	fx.Invoke(func(l *logrus.Logger) {
		
	}),
)

// 同时导出全局访问函数，方便非DI代码调用
func L() *logrus.Logger {
	if logger.GlobalLog == nil {
		panic("Logger 未被初始化")
	}
	return logger.GlobalLog
}
