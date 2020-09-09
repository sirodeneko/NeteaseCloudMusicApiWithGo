package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 通知-私信
func MsgPrivate(c *gin.Context) {
	var service service.MsgPrivateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MsgPrivate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
