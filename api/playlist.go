package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 更新歌单
func PlaylistUpdate(c *gin.Context) {
	var service service.PlaylistUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 更新歌单描述
func PlaylistDescUpdate(c *gin.Context) {
	var service service.PlaylistDescUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistDescUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 更新歌单名
func PlaylistNameUpdate(c *gin.Context) {
	var service service.PlaylistNameUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistNameUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 更新歌单标签
func PlaylistTagsUpdate(c *gin.Context) {
	var service service.PlaylistTagsUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistTagsUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 调整歌单顺序
func PlaylistOrderUpdate(c *gin.Context) {
	var service service.PlaylistOrderUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistOrderUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
