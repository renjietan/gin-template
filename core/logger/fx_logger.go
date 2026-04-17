package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type FxLogger struct {
	Logger *logrus.Logger
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	//fmt.Println(fx_module.L())
	switch e := event.(type) {
	//case *fxevent.Provided:
	//	l.Logger.Info("fx-provided:", e.ConstructorName, "type:", e.OutputTypeNames)
	//case *fxevent.Invoked:
	//if e.Err != nil {
	//	l.Logger.Error("fx-invoke", "function:", e.FunctionName, "error:", e.Err)
	//}
	//case *fxevent.Started: // 应用启动完成
	//	l.Logger.Info("======================= fx应用启动完成 ========================", e)
	//case *fxevent.OnStartExecuting: // 执行 OnStart 钩子前
	//	l.Logger.Info("======================= fx正在启动 ========================")
	case *fxevent.OnStartExecuted: // 执行 OnStart 钩子后
		l.Logger.Info("======================= fx onStart 后 ========================", e.FunctionName, e.Method)
	}
}
