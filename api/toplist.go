package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取所有榜单
func Toplist(c *gin.Context) {
	var service service.ToplistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Toplist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取所有绑定内容摘要
func ToplistDetail(c *gin.Context) {
	var service service.ToplistDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ToplistDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌手榜
func ToplistArtist(c *gin.Context) {
	var service service.ToplistArtistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ToplistArtist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
