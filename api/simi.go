package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func SimiArtist(c *gin.Context) {
	var service service.SimiArtistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SimiArtist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 相似歌单
func SimiPlaylist(c *gin.Context) {
	var service service.SimiPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SimiPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 相似MV
func SimiMv(c *gin.Context) {
	var service service.SimiMvService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SimiMv(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 相似歌手
func SimiSong(c *gin.Context) {
	var service service.SimiSongService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SimiSong(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取最近 5 个听了这首歌的用户
func SimiUser(c *gin.Context) {
	var service service.SimiUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SimiUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
