package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func Like(c *gin.Context) {
	var service service.LikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Like(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
