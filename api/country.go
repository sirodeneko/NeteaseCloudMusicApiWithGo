package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 国家编码列表
func CountriesCodeList(c *gin.Context) {
	var service service.CountriesCodeListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CountriesCodeList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
