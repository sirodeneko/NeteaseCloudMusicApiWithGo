package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func RegisterCellphone(c *gin.Context) {
	var service service.RegisterCellphoneService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RegisterCellphone(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
