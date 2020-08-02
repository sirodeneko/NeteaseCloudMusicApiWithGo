package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func ResourceLike(c *gin.Context) {
	var service service.ResourceLikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResourceLike(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
