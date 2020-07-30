package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SearchDefaultService struct {
}

func (service *SearchDefaultService) SearchDefault(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/search/defaultkeyword/get",
	}
	data := make(map[string]string)

	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/search/defaultkeyword/get`, data, options)

	return reBody
}
