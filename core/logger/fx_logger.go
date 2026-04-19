package logger

import (
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type FxLogger struct {
	Logger        *logrus.Logger
	info          map[string]interface{} // 需要打印的信息
	lock          sync.Mutex
	lifeCycleMame string //声明周期名称
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.info = map[string]interface{}{}
	switch e := event.(type) {
	case *fxevent.Provided:
		if e.Err != nil {
			l.Logger.Error("fxevent.Provided 错误信息: ", e.Err)
			return
		}
		l.info["OutputTypeNames"] = e.OutputTypeNames
		l.info["ConstructorName"] = e.ConstructorName
		l.info["ModuleName"] = e.ModuleName
		l.info["ModuleTrace"] = e.ModuleTrace
		l.info["Private"] = e.Private
		l.info["StackTrace"] = e.StackTrace
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Logger.Error("fxevent.Invoked 错误信息: ", e.Err)
			return
		}
		l.info["FunctionName"] = e.FunctionName
		l.info["ModuleName"] = e.ModuleName
		l.info["Trace"] = e.Trace
	case *fxevent.Started: // 应用启动完成
		if e.Err != nil {
			l.Logger.Error("声明周期（应用启动完成）报错：", e.Err)
			return
		}
	case *fxevent.OnStartExecuting: // 执行 OnStart 钩子前
		l.info["CallerName"] = e.CallerName
		l.info["FunctionName"] = e.FunctionName
	case *fxevent.OnStartExecuted: // 执行 OnStart 钩子后
		if e.Err != nil {
			l.Logger.Error("声明周期（执行 OnStart 钩子后）报错：", e.Err)
			return
		}
		l.info["Runtime"] = e.Runtime
		l.info["CallerName"] = e.CallerName
		l.info["FunctionName"] = e.FunctionName
		l.info["Method"] = e.Method
	}
	l.info["package"] = "fx_logger"
	// 通过反射 获取 event 的名称，最终得到 英文字符
	val := reflect.TypeOf(event)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	l.info["LifeCycleName"] = val.Name()
	l.Logger.WithFields(l.info).Info(l.lifeCycleMame)
}

func (l *FxLogger) Info(msg string) {
	l.Logger.Info(msg)
}
