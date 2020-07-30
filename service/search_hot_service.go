package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SearchHotService struct {
}

func (service *SearchHotService) SearchHot(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
		Ua:      "mobile",
	}
	data := make(map[string]string)
	data["type"] = "1111"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/search/hot`, data, options)

	return reBody
}
