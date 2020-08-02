package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取专辑信息
func Album(c *gin.Context) {
	var service service.AlbumService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Album(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取专辑动态信息
func AlbumDetailDynamic(c *gin.Context) {
	var service service.AlbumDetailDynamicService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumDetailDynamic(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 收藏/取消收藏专辑
func AlbumSub(c *gin.Context) {
	var service service.AlbumSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取已收藏专辑
func AlbumSublist(c *gin.Context) {
	var service service.AlbumSublistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumSublist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
