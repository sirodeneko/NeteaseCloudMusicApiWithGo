package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func FmTrash(c *gin.Context) {
	var service service.FmTrashService
	if err := c.ShouldBind(&service); err == nil {
		res := service.FmTrash(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
