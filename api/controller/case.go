package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Case struct {
}

// case1 godoc
//
//	@Summary	测试接口1
//	@Tags		case
//	@Produce	json
//	@Success	200	{object}	WsClientCountResponse
//	@Router		/case/one [post]
func CaseOne(c *gin.Engine, db *gorm.DB) {
	c.POST("/case/one", func(c *gin.Context) {
		_, err := db.Where("")
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "case1",
		})
	})
}
