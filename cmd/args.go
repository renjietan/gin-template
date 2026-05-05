package cmd

type Args struct {
	/** 注意：
		1、short 标签：只可以是单个字母
		2、短命令使用: -d x=xx -c x=xx
		3、长命令使用: --d x=xx--t x=xx
	**/
	All       bool               `short:"a" long:"all" description:"启动所有服务" required:"false"`
	Swagger   bool               `short:"s" long:"swagger" description:"启动swagger(示例：-s | --s)" required:"false"`
	Driver    map[DriverKey]bool `short:"d" long:"driver" description:"需要启动哪些引擎: (-d redis:true -d mysql:true -d sqlite:true)" required:"false"`
	Ns        bool               `short:"n" long:"ns" description:"是否启动 ns(示例map参数: -n | --ns)" required:"false"`
	Websocket bool               `short:"w" long:"websocket" description:"是否启动 websocket(示例map参数: -w | --websocket)" required:"false"`
	Cron      bool               `short:"c" long:"cron" description:"是否启动 cron(示例map参数: -c | --cron)" required:"false"`
}
