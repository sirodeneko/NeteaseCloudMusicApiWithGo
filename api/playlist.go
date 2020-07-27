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
