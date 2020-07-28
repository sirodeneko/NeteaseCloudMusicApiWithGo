package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 分享xxx到动态
func ShareResource(c *gin.Context) {
	var service service.ShareResourceService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShareResource(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
