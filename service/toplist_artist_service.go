package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type ToplistArtistService struct {
	Type   string `json:"type" form:"type"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *ToplistArtistService) ToplistArtist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Type == "" {
		data["type"] = "1"
	} else {
		data["type"] = service.Type
	}
	if service.Limit == "" {
		data["limit"] = "100"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/toplist/artist`, data, options)

	return reBody
}
