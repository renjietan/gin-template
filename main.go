package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"example.com/t/api/service"
	"example.com/t/core"
	"example.com/t/core/fx_module"
	"example.com/t/core/logger"
	"example.com/t/types"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// AppLifecycle 应用程序生命周期
type AppLifecycle struct {
}

// OnStart 应用程序启动时执行
func (l *AppLifecycle) OnStart(context.Context) error {
	log.Printf("监听服务启动:")
	return nil
}

// OnStop 应用程序停止时执行
func (l *AppLifecycle) OnStop(context.Context) error {
	log.Printf("监听服务停止")
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
				log.Fatal("生产环境 抛出异常(main.go): ", err)
			}
		}()
	}
	app := fx.New(
		// 日志初始化
		fx.Provide(logger.NewLogger),
		// 注入 FX 日志管理器，转为 FX 服务
		//fx.Provide(logger.NewFxLogger),
		fx.WithLogger(func(l *logrus.Logger) fxevent.Logger {
			return &logger.FxLogger{
				Logger: l,
			}
			//return fxevent.NopLogger
		}),
		// 初始化应用配置
		fx.Provide(func() *types.AppConfig {
			config, err := core.LoadConfig(configFile, debug)
			if err != nil {
				log.Fatal(err)
			}
			config.Path = configFile
			if debug {
				_ = core.SaveConfig(config)
			}
			return config
		}),

		// 开启 http-server
		fx_module.GinModule(debug),

		// 初始化数据库
		fx.Provide(core.NewGormConfig),
		fx.Provide(core.NewMysql),
		// 生命周期
		fx.Provide(NewAppLifeCycle),
		// 注册生命周期回调函数
		fx.Invoke(func(lifecycle fx.Lifecycle, lc *AppLifecycle, server *core.AppServer) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return lc.OnStart(ctx)
				},
				OnStop: func(ctx context.Context) error {
					//err := server.Server.Shutdown(ctx)
					//if err != nil {
					//	return err
					//}
					return lc.OnStop(ctx)
				},
			})
		}),
		// 自动同步 数据库 表
		fx.Provide(service.NewMigrationService),
		fx.Invoke(func(migrationService *service.MigrationService) {
			migrationService.StartMigrate()
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
