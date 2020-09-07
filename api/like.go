package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 喜欢一首歌曲
func Like(c *gin.Context) {
	var service service.LikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Like(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取喜欢歌曲列表
func LikeList(c *gin.Context) {
	var service service.LikeListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LikeList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
