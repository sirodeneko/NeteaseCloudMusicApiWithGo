package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func HistoryRecommendSongs(c *gin.Context) {
	var service service.HistoryRecommendSongsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HistoryRecommendSongs(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 历史推荐详情
func HistoryRecommendDongsDetail(c *gin.Context) {
	var service service.HistoryRecommendDongsDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HistoryRecommendDongsDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
