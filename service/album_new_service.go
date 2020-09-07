package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type AlbumNewService struct {
	Area   string `json:"area" form:"area"` //ALL:全部,ZH:华语,EA:欧美,KR:韩国,JP:日本
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *AlbumNewService) AlbumNew(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Area == "" {
		service.Area = "ALL"
	}
	if service.Limit == "" {
		service.Limit = "30"
	}
	if service.Offset == "" {
		service.Offset = "0"
	}
	data["area"] = service.Area
	data["limit"] = service.Limit
	data["offset"] = service.Offset
	data["total"] = "true"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/album/new`, data, options)

	return reBody
}
