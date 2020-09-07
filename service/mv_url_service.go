package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type MvUrlService struct {
	ID string `json:"id" form:"id"`
	R  string `json:"r" form:"r"`
}

func (service *MvUrlService) MvUrl(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	if service.R == "" {
		data["r"] = "1080"
	} else {
		data["r"] = service.R
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/song/enhance/play/mv/url`, data, options)

	return reBody
}
