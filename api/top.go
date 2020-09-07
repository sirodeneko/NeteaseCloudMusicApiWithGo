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

// 新歌速递
func TopSong(c *gin.Context) {
	var service service.TopSongService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopSong(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 新碟上架
func TopAlbum(c *gin.Context) {
	var service service.TopAlbumService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopAlbum(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门歌手
func TopArtists(c *gin.Context) {
	var service service.TopArtistsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopArtists(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// mv排行
func TopMv(c *gin.Context) {
	var service service.TopMvService
	if err := c.ShouldBind(&service); err == nil {
		res := service.TopMv(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
