package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 搜索
func Search(c *gin.Context) {
	var service service.SearchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Search(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 默认搜索关键词
func SearchDefault(c *gin.Context) {
	var service service.SearchDefaultService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchDefault(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门搜索关键词
func SearchHot(c *gin.Context) {
	var service service.SearchHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热搜列表（详细）
func SearchHotDetail(c *gin.Context) {
	var service service.SearchHotDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchHotDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 搜索建议
func SearchSuggest(c *gin.Context) {
	var service service.SearchSuggestService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchSuggest(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 多重搜索
func SearchMultimatch(c *gin.Context) {
	var service service.SearchMultimatchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchMultimatch(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
