package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type HomepageBlockPageService struct {
	Refresh string `json:"refresh" form:"refresh"`
}

func (service *HomepageBlockPageService) HomepageBlockPage(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Refresh == "" {
		service.Refresh = "true"
	}
	data["refresh"] = service.Refresh
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/homepage/block/page`, data, options)

	return reBody
}
