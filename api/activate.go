package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 注册后初始化昵称
func ActivateInitProfile(c *gin.Context) {
	var service service.ActivateInitProfileService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ActivateInitProfile(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
