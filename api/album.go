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
