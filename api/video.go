package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func VideoSub(c *gin.Context) {
	var service service.VideoSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
