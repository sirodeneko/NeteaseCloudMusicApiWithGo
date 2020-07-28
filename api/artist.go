package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func ArtistList(c *gin.Context) {
	var service service.ArtistListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArtistList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
