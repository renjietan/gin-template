package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"example.com/t/api/controller"
	"example.com/t/core"
	"example.com/t/logger"
	"example.com/t/types"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var log2 = logger.GetLogger()

// AppLifecycle 应用程序生命周期
type AppLifecycle struct {
}

// OnStart 应用程序启动时执行
func (l *AppLifecycle) OnStart(context.Context) error {
	log2.Info("监听服务启动")
	return nil
}

// OnStop 应用程序停止时执行
func (l *AppLifecycle) OnStop(context.Context) error {
	log2.Info("监听服务停止")
	return nil
}

func NewAppLifeCycle() *AppLifecycle {
	return &AppLifecycle{}
}

func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.toml"
	}
	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if !debug {
		defer func() {
			if err := recover(); err != nil {
				log2.Error("生产环境 抛出异常（main.go）：:", err)
			}
		}()
	}

	app := fx.New(
		fx.Provide(func() *zap.SugaredLogger { return log2 }),
		// 将 fx 日志 统一使用日志管理器收集
		fx.WithLogger(func(sugar *zap.SugaredLogger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: sugar.Desugar(),
			}
		}),
		// 初始化配置应用配置
		fx.Provide(func() *types.AppConfig {
			config, err := core.LoadConfig(configFile)
			if err != nil {
				log2.Fatal(err)
			}
			config.Path = configFile
			if debug {
				_ = core.SaveConfig(config)
			}
			return config
		}),
		fx.Provide(core.NewAppServer),
		fx.Invoke(func(appserver *core.AppServer) {
			appserver.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}),
		fx.Invoke(func(appserver *core.AppServer) {
			appserver.Run()
			appserver.Middlewares(debug)
			controller.Case1(appserver.Engine)
		}),
		fx.Provide(NewAppLifeCycle),
		// 注册生命周期回调函数
		fx.Invoke(func(lifecycle fx.Lifecycle, lc *AppLifecycle, server *core.AppServer) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return lc.OnStart(ctx)
				},
				OnStop: func(ctx context.Context) error {
					err := server.Server.Shutdown(ctx)
					if err != nil {
						return err
					}
					return lc.OnStop(ctx)
				},
			})
		}),
	)

	// 启动应用程序
	go func() {
		if err := app.Start(context.Background()); err != nil {
			log.Fatal("服务启动失败：", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 关闭应用程序
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(ctx); err != nil {
		log.Fatal("任务停止失败（main.go）:", err)
	}
}
