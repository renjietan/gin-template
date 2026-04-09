package controller

import "github.com/gin-gonic/gin"

// case1 godoc
//
//	@Summary	测试接口1
//	@Tags		case
//	@Produce	json
//	@Success	200	{object}	WsClientCountResponse
//	@Router		/case/1 [post]
func Case1(c *gin.Engine) {
	c.POST("/case/1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "case1",
		})
	})
}
