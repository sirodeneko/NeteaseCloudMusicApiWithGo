package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 歌单推荐
func RelatedPlaylist(c *gin.Context) {
	var service service.RelatedPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RelatedPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 相关视频
func RelatedAllVideo(c *gin.Context) {
	var service service.RelatedAllVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RelatedAllVideo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
