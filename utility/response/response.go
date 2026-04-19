package response

import (
	"net/http"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
)

func SUCCESS(c *gin.Context, values ...interface{}) {
	if values != nil {
		c.JSON(http.StatusOK, types.Response{Code: types.Success, Data: values[0]})
	} else {
		c.JSON(http.StatusOK, types.Response{Code: types.Success})
	}

}

func ERROR(c *gin.Context, messages ...string) {
	if messages != nil {
		c.JSON(http.StatusBadRequest, types.Response{Code: types.Failed, Message: messages[0]})
	} else {
		c.JSON(http.StatusBadRequest, types.Response{Code: types.Failed})
	}
}

func HACKER(c *gin.Context) {
	c.JSON(http.StatusBadRequest, types.Response{Code: types.Failed, Message: "Hacker attempt!!!"})
}

func NotAuth(c *gin.Context, messages ...string) {
	if messages != nil {
		c.JSON(http.StatusUnauthorized, types.Response{Code: types.NotAuthorized, Message: messages[0]})
	} else {
		c.JSON(http.StatusUnauthorized, types.Response{Code: types.NotAuthorized, Message: "Not Authorized"})
	}
}
