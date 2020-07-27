package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CellphoneExistenceCheck(c *gin.Context) {
	var service service.CellphoneExistenceCheckService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CellphoneExistenceCheck(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
