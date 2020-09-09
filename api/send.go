package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func SendText(c *gin.Context) {
	var service service.SendTextService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendText(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 发生带歌单的私信
func SendPlaylist(c *gin.Context) {
	var service service.SendPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
