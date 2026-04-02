package main

import (
	"net"
)

func main() {
	// 监听TCP端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			continue // 继续接受其他连接
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// 处理连接逻辑
}
