package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 调整歌曲顺序
func SongOrderUpdate(c *gin.Context) {
	var service service.SongOrderUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SongOrderUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取音乐的url
func SongUrl(c *gin.Context) {
	var service service.SongUrlService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SongUrl(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取歌曲详情
func SongDetail(c *gin.Context) {
	var service service.SongDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SongDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
