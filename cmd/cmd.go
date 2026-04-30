package cmd

import (
	"fmt"
	"os"

	"example.com/t/core/fx_module"
	"github.com/jessevdk/go-flags"
	"go.uber.org/fx"
)

func InitFxModule() fx.Option {
	var args Args
	// 使用 flags.Parse 解析并自动进行必填项验证
	if _, err := flags.Parse(&args); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrRequired {
			fmt.Println("\n❌ 错误：缺少必需的参数！")
			// 提示用户使用帮助指令
			fmt.Println("❌ 请使用 -h 或 --help 查看所有必需的选项。\n")
			os.Exit(0)
		} else {
			fmt.Printf("❌ 解析错误: %v\n", err)
		}
		os.Exit(1)
	}
	fmt.Println("✅ 参数验证通过...")
	options := appendModule(&args)
	return fx.Options(options...)
}

func appendModule(args *Args) []fx.Option {
	var fxOptions []fx.Option
	fxOptions = append(fxOptions)
	if args.All {
		fxOptions = append(fxOptions, fx_module.FXNacosModule)
		fxOptions = append(fxOptions, fx_module.FXSQLiteModule)
		fxOptions = append(fxOptions, fx_module.FXMySqlModule)
		fxOptions = append(fxOptions, fx_module.FxRedisModule)
		fxOptions = append(fxOptions, fx_module.FXCronModule)
		fxOptions = append(fxOptions, fx_module.FxWsModule)
		return fxOptions
	}
	if args.Ns {
		fxOptions = append(fxOptions, fx_module.FXNacosModule)
	}
	if args.Driver["sqlite"] {
		fxOptions = append(fxOptions, fx_module.FXSQLiteModule)
	}
	if args.Driver["mysql"] {
		fxOptions = append(fxOptions, fx_module.FXMySqlModule)
	}
	if args.Driver["redis"] {
		fxOptions = append(fxOptions, fx_module.FxRedisModule)
	}
	if args.Cron {
		fxOptions = append(fxOptions, fx_module.FXCronModule)
	}
	if args.Websocket {
		fxOptions = append(fxOptions, fx_module.FxWsModule)
	}
	return fxOptions
}
