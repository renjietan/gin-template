package cmd

import (
	"fmt"
	"os/exec"

	logger2 "example.com/t/logger"
)

func Command_swag() {
	// 构建命令
	cmd := exec.Command("swag", "init")

	// 设置命令执行的工作目录
	cmd.Dir = "." // 替换为你的项目路径

	// 设置命令的标准输出和错误输出
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	cmd.Stdout = logger2.GetLogWriter()
	cmd.Stderr = logger2.GetLogWriter()
	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}
