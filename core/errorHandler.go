package core

import (
	"net/http"
	"runtime/debug"

	"example.com/t/types"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

func errorHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("errorHandler 异常捕获: %v", r)
			debug.PrintStack()
			c.JSON(http.StatusBadRequest, types.Response{Code: types.Failed, Message: types.ErrorMsg})
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
