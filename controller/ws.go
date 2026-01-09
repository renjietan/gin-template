package controller

import (
	"github.com/gin-gonic/gin"

	"example.com/t/ws"
)

// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetWsClientCount(c *gin.Context) {
	ws := c.MustGet("ws").(*ws.WebSocketManager)
	count := ws.GetClientCount()
	c.JSON(200, gin.H{
		"client_count": count,
	})
}

func BroadcastWsMessage(c *gin.Context) {
	ws := c.MustGet("ws").(*ws.WebSocketManager)
	var req struct {
		Msg string `json:"msg"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Msg == "" {
		c.JSON(400, gin.H{"error": "需要字段 msg"})
		return
	}

	ws.Broadcast([]byte(req.Msg))
	c.JSON(200, gin.H{"status": "ok"})
}
