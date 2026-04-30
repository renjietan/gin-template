package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"example.com/t/cmd"
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
				log.Fatal("❌ 生产环境 抛出异常(main.go): ", err)
			}
		}()
	}
	options := cmd.InitFxModule()
	app := fx.New(
		// 日志初始化
		fx.Provide(logger.NewLogger),
		fx.WithLogger(func(l *logrus.Logger) fxevent.Logger {
			// TODO: 此处可能还需优化
			return &logger.FxLogger{
				Logger: l,
			}
			//return fxevent.NopLogger
		}),
		// 初始化应用配置
		fx.Provide(func() *types.AppConfig {
			config, err := core.LoadConfig(configFile)
			if err != nil {
				// 此处无法 使用 logrus
				// 因为 logger.NewLogger 中引入了 AppConfig， 此处 引入 logrus  会导致循环依赖引入
				log.Fatal("❌ 配置文件：读取失败")
			}
			config.Path = configFile
			config.Debug = debug
			if debug {
				_ = core.SaveConfig(config)
			}
			return config
		}),

		// 开启 http-server
		fx_module.FxGinModule,
		fx_module.FxGormConfigModule,
		//// 1、mysql
		//fx_module.FXMySqlModule,
		//// 2、sqlite
		//fx_module.FXSQLiteModule,
		//// 3、redis
		//fx_module.FxRedisModule,
		//// 4、nacos
		//fx_module.FXNacosModule,
		//// 5、websocket
		//fx_module.FxWsModule,
		//// 6、cron
		//fx_module.FXCronModule,

		options,
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
		logger.L().Fatal(err)
	}
}
