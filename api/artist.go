package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取歌手列表
func ArtistList(c *gin.Context) {
	var service service.ArtistListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArtistList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 收藏/取消收藏歌手
func ArtistSub(c *gin.Context) {
	var service service.ArtistSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArtistSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取歌手热门50曲
func ArtistTopSong(c *gin.Context) {
	var service service.ArtistTopSongService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArtistTopSong(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取收藏的歌手
func ArtistSublist(c *gin.Context) {
	var service service.ArtistSublistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ArtistSublist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
