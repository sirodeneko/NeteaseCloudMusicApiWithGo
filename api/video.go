package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func VideoSub(c *gin.Context) {
	var service service.VideoSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取视频分类下的相关视频
func VideoGroup(c *gin.Context) {
	var service service.VideoGroupService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoGroup(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取全部视频列表
func VideoTimelineAll(c *gin.Context) {
	var service service.VideoTimelineAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoTimelineAll(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取推荐视频
func VideoTimelineRecommend(c *gin.Context) {
	var service service.VideoTimelineRecommendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoTimelineRecommend(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
