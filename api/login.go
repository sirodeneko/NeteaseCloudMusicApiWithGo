package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取登录状态
func LoginStatus(c *gin.Context) {
	var service service.LoginStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginStatus(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
