package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 歌单 ( 网友精选碟 )
func TopPlaylist(c *gin.Context) {
	var service service.TopPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 精选歌单
func TopPlaylistHighquality(c *gin.Context) {
	var service service.TopPlaylistHighqualityService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopPlaylistHighquality(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
