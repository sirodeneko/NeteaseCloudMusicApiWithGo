package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SearchSuggestService struct {
	S    string `json:"keywords" form:"keywords"`
	Type string `json:"type" form:"type"`
}

func (service *SearchSuggestService) SearchSuggest(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Type == "mobile" {
		service.Type = "keyword"
	} else {
		service.Type = "web"
	}
	data["s"] = service.S
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/search/suggest/`+service.Type, data, options)

	return reBody
}
