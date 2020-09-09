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

// 私信内容
func MsgPrivateHistory(c *gin.Context) {
	var service service.MsgPrivateHistoryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MsgPrivateHistory(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 通知-评论
func MsgComments(c *gin.Context) {
	var service service.MsgCommentsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MsgComments(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
