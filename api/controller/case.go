package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Case struct {
	id   string
	name string
}

// case1 godoc
//
//	@Summary	测试接口4444
//	@Tags		case
//	@Produce	json
//	@Success	200	{object}	WsClientCountResponse
//	@Router		/case/one1 [post]
func CaseOne(c *gin.Engine, db *gorm.DB) {
	c.POST("/case/one1", func(c *gin.Context) {
		// db.Create(&model.Config{
		// 	Value: "value",
		// 	Name:  "name",
		// })
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "case1",
		})
	})
}
