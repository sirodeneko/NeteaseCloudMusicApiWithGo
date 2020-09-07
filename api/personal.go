package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 私人fm
func PersonalFm(c *gin.Context) {
	var service service.PersonalFmService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalFm(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 推荐mv
func PersonalizedMv(c *gin.Context) {
	var service service.PersonalizedMvService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalizedMv(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 推荐歌单
func Personalized(c *gin.Context) {
	var service service.PersonalizedService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Personalized(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 推荐新音乐
func PersonalizedNewsong(c *gin.Context) {
	var service service.PersonalizedNewsongService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalizedNewsong(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 推荐电台
func PersonalizedDjprogram(c *gin.Context) {
	var service service.PersonalizedDjprogramService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalizedDjprogram(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 独家放送（入口列表）
func PersonalizedPrivatecontent(c *gin.Context) {
	var service service.PersonalizedPrivatecontentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalizedPrivatecontent(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 独家放送列表
func PersonalizedPrivatecontentList(c *gin.Context) {
	var service service.PersonalizedPrivatecontentListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.PersonalizedPrivatecontentList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
