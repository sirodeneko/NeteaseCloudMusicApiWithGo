package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CommentEvent(c *gin.Context) {
	var service service.CommentEventService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentEvent(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 云村热评
func CommentHotwallList(c *gin.Context) {
	var service service.CommentHotwallListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentHotwallList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
