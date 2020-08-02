package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 私人fm
func PersonalFm(c *gin.Context) {
	var service service.PersonalFmService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalFm(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
