package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func Banner(c *gin.Context) {
	var service service.BannerService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Banner(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
