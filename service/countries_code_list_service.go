package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type CountriesCodeListService struct {
}

func (service *CountriesCodeListService) CountriesCodeList(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/lbs/countries/v1",
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/lbs/countries/v1`, data, options)

	return reBody
}
