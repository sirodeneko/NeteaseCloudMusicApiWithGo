package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 签到
func DailySignin(c *gin.Context) {
	var service service.DailySigninService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DailySignin(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
