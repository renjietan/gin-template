package logger

import (
	"fmt"
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
		l.info["OutputTypeNames"] = e.OutputTypeNames
		l.info["ConstructorName"] = e.ConstructorName
		l.info["ModuleName"] = e.ModuleName
		l.info["ModuleTrace"] = e.ModuleTrace
		l.info["Private"] = e.Private
		l.info["StackTrace"] = e.StackTrace
		if e.Err != nil {
			l.Logger.Error("fxevent.Provided 错误信息: ", e.Err)
			return
		}
	case *fxevent.Invoked:
		l.info["FunctionName"] = e.FunctionName
		l.info["ModuleName"] = e.ModuleName
		l.info["Trace"] = e.Trace
		if e.Err != nil {
			l.Logger.Error("fxevent.Invoked 错误信息: ", e.Err)
			return
		}
	case *fxevent.Started: // 应用启动完成
		if e.Err != nil {
			l.Logger.Error("声明周期（应用启动完成）报错：", e.Err)
			return
		}
	case *fxevent.OnStartExecuting: // 执行 OnStart 钩子前
		l.info["CallerName"] = e.CallerName
		l.info["FunctionName"] = e.FunctionName
	case *fxevent.OnStartExecuted: // 执行 OnStart 钩子后
		l.info["Runtime"] = e.Runtime
		l.info["CallerName"] = e.CallerName
		l.info["FunctionName"] = e.FunctionName
		l.info["Method"] = e.Method
		if e.Err != nil {
			l.Logger.Error("声明周期（执行 OnStart 钩子后）报错：", e.Err)
			return
		}
	}

	val := reflect.TypeOf(event)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	msg := BeautifyJsonStr(l.info)
	name := fmt.Sprintf("✅ %s", val.Name())
	l.Logger.WithFields(logrus.Fields{
		"package": "fx_logger",
		"msg":     msg,
	}).Info(name)
}

func (l *FxLogger) info2string(event fxevent.Event) (msg string, name string) {
	val := reflect.TypeOf(event)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	msg = BeautifyJsonStr(l.info)
	name = fmt.Sprintf("✅ %s", val.Name())
	return msg, name
}

func (l *FxLogger) Info(event fxevent.Event) {
	msg, name := l.info2string(event)
	l.Logger.WithFields(logrus.Fields{
		"package": "fx",
		"msg":     msg,
	}).Info(name)
}

func (l *FxLogger) Debug(event fxevent.Event) {
	fmt.Println("====================================== Debug =============================================")
	msg, name := l.info2string(event)
	l.Logger.WithFields(logrus.Fields{
		"package": "fx",
		"msg":     msg,
	}).Debug(name)
}

func (l *FxLogger) Warn(msg string) {
	fmt.Println("====================================== Warn =============================================")
	l.Logger.Warn(msg)
}

func (l *FxLogger) Error(msg string) {
	fmt.Println("====================================== Error =============================================")
	l.Logger.Error(msg)
}

func (l *FxLogger) Fatal(msg string) {
	l.Logger.Fatal(msg)
}

func (l *FxLogger) Panic(msg string) {
	l.Logger.Panic(msg)
}

func (l *FxLogger) Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *FxLogger) Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func (l *FxLogger) Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func (l *FxLogger) Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func (l *FxLogger) Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func (l *FxLogger) Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
