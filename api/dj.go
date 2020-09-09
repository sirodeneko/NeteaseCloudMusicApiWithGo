package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 电台banner
func DjBanner(c *gin.Context) {
	var service service.DjBannerService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjBanner(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门电台
func DjHot(c *gin.Context) {
	var service service.DjHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台节目榜
func DjProgramToplist(c *gin.Context) {
	var service service.DjProgramToplistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjProgramToplist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台付费精品
func DjToplistPay(c *gin.Context) {
	var service service.DjToplistPayService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjToplistPay(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-24小时节目榜
func DjProgramToplistHours(c *gin.Context) {
	var service service.DjProgramToplistHoursService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjProgramToplistHours(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-24小时主播榜
func DjToplistHours(c *gin.Context) {
	var service service.DjToplistHoursService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjToplistHours(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-主播新人榜
func DjToplistNewcomer(c *gin.Context) {
	var service service.DjToplistNewcomerService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjToplistNewcomer(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-最热主播榜
func DjToplistPopular(c *gin.Context) {
	var service service.DjToplistPopularService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjToplistPopular(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台 - 新晋电台榜/热门电台榜
func DjToplist(c *gin.Context) {
	var service service.DjToplistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjToplist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-类别热门电台
func DjRadioHot(c *gin.Context) {
	var service service.DjRadioHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjRadioHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-推荐
func DjRecommend(c *gin.Context) {
	var service service.DjRecommendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjRecommend(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-分类
func DjCatelist(c *gin.Context) {
	var service service.DjCatelistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjCatelist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-分类推荐
func DjRecommendType(c *gin.Context) {
	var service service.DjRecommendTypeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjRecommendType(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-订阅
func DjSub(c *gin.Context) {
	var service service.DjSubService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjSub(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-订阅列表
func DjSublist(c *gin.Context) {
	var service service.DjSublistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjSublist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-付费精选
func DjPaygift(c *gin.Context) {
	var service service.DjPaygiftService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjPaygift(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-非热门类型
func DjCategoryExcludehot(c *gin.Context) {
	var service service.DjCategoryExcludehotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjCategoryExcludehot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台推荐类型
func DjCategoryRecommend(c *gin.Context) {
	var service service.DjCategoryRecommendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjCategoryRecommend(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-今日优选
func DjTodayPerfered(c *gin.Context) {
	var service service.DjTodayPerferedService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjTodayPerfered(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-详情
func DjDetail(c *gin.Context) {
	var service service.DjDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-节目
func DjProgram(c *gin.Context) {
	var service service.DjProgramService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjProgram(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台-节目详情
func DjProgramDetail(c *gin.Context) {
	var service service.DjProgramDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DjProgramDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
