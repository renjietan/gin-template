package cmd

import (
	"fmt"
	"os/exec"
)

func Command_swag() {
	cmd := exec.Command("swag", "init")
	// 设置命令执行的工作目录
	cmd.Dir = "."
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}
