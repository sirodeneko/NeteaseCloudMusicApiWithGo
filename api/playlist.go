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

// 获取歌单分类
func PlaylistCatlist(c *gin.Context) {
	var service service.PlaylistCatlistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistCatlist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门歌单分类
func PlaylistHot(c *gin.Context) {
	var service service.PlaylistHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌单详情
func PlaylistDetail(c *gin.Context) {
	var service service.PlaylistDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 创建歌单
func PlaylistCreate(c *gin.Context) {
	var service service.PlaylistCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistCreate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 删除歌单
func PlaylistDelete(c *gin.Context) {
	var service service.PlaylistDeleteService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistDelete(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 收藏歌单
func PlaylistSubscribe(c *gin.Context) {
	var service service.PlaylistSubscribeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistSubscribe(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌单收藏者
func PlaylistSubscribers(c *gin.Context) {
	var service service.PlaylistSubscribersService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistSubscribers(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 添加或者删除歌曲
func PlaylistTracks(c *gin.Context) {
	var service service.PlaylistTracksService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PlaylistTracks(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
