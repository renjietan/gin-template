package main

import (
	"fmt"
	"os"

	"example.com/t/utility"
	"github.com/jessevdk/go-flags"
)

const (
	Redis  = "redis"
	Sqlite = "sqlite"
	Mysql  = "mysql"
)

type DriverKey string

func (r *DriverKey) UnmarshalFlag(value string) error {
	switch value {
	case Redis, Sqlite, Mysql:
		*r = DriverKey(value)
		return nil
	default:
		return fmt.Errorf("错误的键值: %q, 必须是 %q | %q | %q",
			value, Redis, Sqlite, Mysql)
	}
}

// Options 通过结构体标签来定义所有的命令行参数
type Args struct {
	/** 注意：
		1、short 标签：只可以是单个字母
		2、短命令使用: -d x=xx -c x=xx
		3、长命令使用: --d x=xx--t x=xx
	**/
	All       bool               `short:"a" long:"all" description:"启动所有服务" required:"false"`
	Driver    map[DriverKey]bool `short:"d" long:"driver" description:"需要启动哪些引擎: (-d redis:true -d mysql:true -d sqlite:true )" required:"true"`
	Ns        bool               `short:"n" long:"ns" description:"是否启动 ns(示例map参数: --ns=true)" required:"false"`
	Websocket bool               `short:"w" long:"websocket" description:"是否启动 websocket(示例map参数: --websocket=true)" required:"false"`
	Cron      bool               `short:"c" long:"cron" description:"是否启动 cron(示例map参数: --cron=true)" required:"false"`
}

func main() {
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
	fmt.Println("\n✅ 参数验证通过，程序继续执行...")
	fmt.Printf("参数: %s\n", utility.Interface2String(args))

	var fxOptions []int
	fxOptions = append(fxOptions, 2, 3, 4, 5, 6)
	fmt.Println(fxOptions)
}
