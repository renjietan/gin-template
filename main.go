package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"example.com/t/udp"
	"example.com/t/ws"
	"github.com/gin-gonic/gin"
)

// 心跳间隔
const heartbeatInterval = 3 * time.Second

const serverAddr = "127.0.0.1:9000"

var udpClient *udp.UDPClient
var wsManager *ws.WebSocketManager

func main() {
	// ============ 1. 初始化 Gin ============ //
	r := gin.Default()

	// ============ 2. 初始化 WebSocket ============ //
	wsManager = ws.NewWebSocketManager()
	defer wsManager.Close()

	// WebSocket 路由
	r.GET("/ws", wsManager.HandleWebSocket)

	// 获取 WebSocket 连接数
	r.GET("/ws/count", func(c *gin.Context) {
		count := wsManager.GetClientCount()
		c.JSON(200, gin.H{
			"client_count": count,
		})
	})

	// 广播消息接口
	r.POST("/ws/broadcast", func(c *gin.Context) {
		var req struct {
			Msg string `json:"msg"`
		}
		if err := c.ShouldBindJSON(&req); err != nil || req.Msg == "" {
			c.JSON(400, gin.H{"error": "需要字段 msg"})
			return
		}

		wsManager.Broadcast([]byte(req.Msg))
		c.JSON(200, gin.H{"status": "ok"})
	})

	// ============ 3. UDP 相关路由 ============ //
	r.GET("/udp/last", func(c *gin.Context) {
		if udpClient == nil {
			c.JSON(500, gin.H{"error": "UDP 客户端尚未初始化"})
			return
		}

		msg := udpClient.LastMsg()

		c.JSON(200, gin.H{
			"last_msg": msg,
		})
	})

	r.POST("/udp/send", func(c *gin.Context) {
		var req struct {
			Msg string `json:"msg"`
		}
		if err := c.ShouldBindJSON(&req); err != nil || req.Msg == "" {
			c.JSON(400, gin.H{"error": "需要字段 msg"})
			return
		}

		if udpClient == nil {
			c.JSON(500, gin.H{"error": "UDP 客户端尚未初始化"})
			return
		}

		if err := udpClient.Send(req.Msg); err != nil {
			c.JSON(500, gin.H{"error": "发送失败", "detail": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})

	// ============ 4. 初始化 UDP 客户端 ============ //
	var err error
	udpClient, err = udp.NewUDPClient(serverAddr, heartbeatInterval)
	if err != nil {
		fmt.Println("UDP 客户端初始化失败:", err)
		return
	}
	defer udpClient.Close()

	// 程序退出信号（Ctrl+C）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// ============ 5. 启动 Gin HTTP 服务 ============ //
	go func() {
		if err := r.Run(":8080"); err != nil {
			fmt.Println("Gin 启动失败:", err)
		}
	}()

	// 等待退出信号（Ctrl+C）
	<-quit
	fmt.Println("\n收到退出信号, 准备关闭...")
	// 稍等一会儿保证 goroutine 退出（Close 中会关闭连接并通知协程）
	fmt.Println("客户端已退出")
}
