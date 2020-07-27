package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 发送验证码
func CaptchaSent(c *gin.Context) {
	var service service.CaptchaSentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CaptchaSent(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 验证验证码
func CaptchaVerify(c *gin.Context) {
	var service service.CaptchaVerifyService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CaptchaVerify(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
