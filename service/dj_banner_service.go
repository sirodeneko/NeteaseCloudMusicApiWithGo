package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type DjBannerService struct {
}

func (service *DjBannerService) DjBanner(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/banner/get`, data, options)

	return reBody
}
