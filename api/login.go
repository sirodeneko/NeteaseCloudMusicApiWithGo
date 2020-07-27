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

// 手机登录
func LoginCellphone(c *gin.Context) {
	var service service.LoginCellphoneService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginCellphone(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 邮箱登录
func LoginEmail(c *gin.Context) {
	var service service.LoginEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginEmail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 刷新登录
func LoginRefresh(c *gin.Context) {
	var service service.LoginRefreshService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoginRefresh(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
