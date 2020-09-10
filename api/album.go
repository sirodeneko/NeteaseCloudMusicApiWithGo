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

// 全部新碟
func AlbumNew(c *gin.Context) {
	var service service.AlbumNewService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumNew(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 数字专辑-新碟上架
func AlbumList(c *gin.Context) {
	var service service.AlbumListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 数字专辑&数字单曲-榜单
func AlbumSongsaleboard(c *gin.Context) {
	var service service.AlbumSongsaleboardService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumSongsaleboard(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 数字专辑-语种风格馆
func AlbumListStyle(c *gin.Context) {
	var service service.AlbumListStyleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumListStyle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 数字专辑详情
func AlbumDetail(c *gin.Context) {
	var service service.AlbumDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AlbumDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 我的数字专辑
func DigitalAlbumPurchased(c *gin.Context) {
	var service service.DigitalAlbumPurchasedService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DigitalAlbumPurchased(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 购买数字专辑
func DigitalAlbumOrdering(c *gin.Context) {
	var service service.DigitalAlbumOrderingService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DigitalAlbumOrdering(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
