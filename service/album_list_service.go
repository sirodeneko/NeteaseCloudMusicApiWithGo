package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type AlbumListService struct {
	Area   string `json:"area" form:"area"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
	Type   string `json:"type" form:"type"`
}

func (service *AlbumListService) AlbumList(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.Area == "" {
		data["area"] = "ALL"
	} else {
		data["area"] = service.Offset
	}
	data["order"] = "true"
	data["type"] = service.Type
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/vipmall/albumproduct/list`, data, options)

	return reBody
}
