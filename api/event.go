package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 转发动态
func EventForward(c *gin.Context) {
	var service service.EventForwardService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EventForward(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 删除动态
func EventDel(c *gin.Context) {
	var service service.EventDelService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EventDel(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
