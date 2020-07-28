package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 获取用户信息
func UserDetail(c *gin.Context) {
	var service service.UserDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户信息 , 歌单，收藏，mv, dj 数量
func UserSubcount(c *gin.Context) {
	var service service.UserSubcountService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserSubcount(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 更新用户信息
func UserUpdate(c *gin.Context) {
	var service service.UserUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserUpdate(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户歌单
func UserPlaylist(c *gin.Context) {
	var service service.UserPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户电台
func UserDj(c *gin.Context) {
	var service service.UserDjService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserDj(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户关注列表
func UserFollows(c *gin.Context) {
	var service service.UserFollowsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserFollows(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户粉丝列表
func UserFolloweds(c *gin.Context) {
	var service service.UserFollowedsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserFolloweds(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取用户动态
func UserEvent(c *gin.Context) {
	var service service.UserEventService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserEvent(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
