package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func Scrobble(c *gin.Context) {
	var service service.ScrobbleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Scrobble(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
