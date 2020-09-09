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

// 视频标签列表
func VideoGroupList(c *gin.Context) {
	var service service.VideoGroupListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoGroupList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频分类列表
func VideoCategoryList(c *gin.Context) {
	var service service.VideoCategoryListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoCategoryList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频详情
func VideoDetail(c *gin.Context) {
	var service service.VideoDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取视频点赞转发评论数数据
func VideoDetailInfo(c *gin.Context) {
	var service service.VideoDetailInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoDetailInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取视频播放地址
func VideoUrl(c *gin.Context) {
	var service service.VideoUrlService
	if err := c.ShouldBind(&service); err == nil {
		res := service.VideoUrl(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
