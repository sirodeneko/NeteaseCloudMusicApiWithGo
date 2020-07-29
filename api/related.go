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
