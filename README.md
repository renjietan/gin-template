// @title           xxx API
// @version         1.0
// @description     xxx统 API 服务
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
//// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
"fmt"
"os"
"os/signal"
"syscall"
"time"

	"example.com/t/controller"
	"example.com/t/udp"
	"example.com/t/ws"
	"github.com/gin-gonic/gin" // 新版需要导入 renderer 包
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 心跳间隔
const heartbeatInterval = 3 * time.Second

const serverAddr = "8.135.10.183:53753"

var udpClient *udp.UDPClient
var wsManager *ws.WebSocketManager

// UDPMessageResponse 用于 swagger 展示最后一条 UDP 消息.
type UDPMessageResponse struct {
LastMsg string `json:"last_msg" example:"{...}"`
}

// SendUDPRequest 用于 swagger 展示发送消息体.
type SendUDPRequest struct {
Msg string `json:"msg" example:"hello udp"`
}

func main() {
// Swagger 基础路径
// docs.SwaggerInfo.BasePath = "/"
// ============ 1. 初始化 Gin ============ //
r := gin.Default()
// ============ 2. 初始化 WebSocket ============ //
wsManager = ws.NewWebSocketManager()
r.Use(func(c *gin.Context) {
c.Set("ws", wsManager)
c.Next()
})

	defer wsManager.Close()

	// WebSocket 路由
	if gin.Mode() != gin.ReleaseMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ws", wsManager.HandleWebSocket)

	// 获取 WebSocket 连接数
	r.GET("/ws/count", controller.GetWsClientCount)

	// 广播消息接口
	r.POST("/ws/broadcast", controller.BroadcastWsMessage)

	// ============ 3. UDP 相关路由 ============ //
	r.GET("/udp/last", getLastUDP)

	r.POST("/udp/send", sendUDP)

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
		if gin.Mode() != gin.ReleaseMode {
			// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler),
			// 	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
			// 	ginSwagger.DefaultModelsExpandDepth(-1),
			// )
			// ginSwagger.WrapHandler(swaggerFiles.Handler)
			if err := r.Run(":8080"); err != nil {
				fmt.Println("Gin 启动失败:", err)
			}
		}
	}()

	<-quit
	fmt.Println("\n收到退出信号, 准备关闭...")
	fmt.Println("客户端已退出")
}

// getLastUDP godoc
// @Summary 获取最后一条 UDP 消息
// @Tags udp
// @Produce json
// @Success 200 {object} UDPMessageResponse
// @Failure 500 {object} controller.ErrorResponse
// @Router /udp/last [get]
func getLastUDP(c *gin.Context) {
if udpClient == nil {
c.JSON(500, gin.H{"error": "UDP 客户端尚未初始化"})
return
}

	msg := udpClient.LastMsg()

	c.JSON(200, gin.H{
		"last_msg": msg,
	})
}

// sendUDP godoc
// @Summary 发送 UDP 消息
// @Tags udp
// @Accept JSON
// @Produce JSON
// @Param data body SendUDPRequest true "消息体"
// @Success 200 {object} controller.StatusResponse
// @Failure 400 {object} controller.ErrorResponse
// @Failure 500 {object} controller.ErrorResponse
// @Router /udp/send [post]
func sendUDP(c *gin.Context) {
var req SendUDPRequest
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
}
