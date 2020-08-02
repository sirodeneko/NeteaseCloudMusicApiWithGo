package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 日推歌单
func RecommendResource(c *gin.Context) {
	var service service.RecommendResourceService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RecommendResource(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 日推歌曲
func RecommendSongs(c *gin.Context) {
	var service service.RecommendSongsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RecommendSongs(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
