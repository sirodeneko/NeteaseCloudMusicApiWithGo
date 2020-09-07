package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type PersonalizedService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *PersonalizedService) Personalized(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

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
	data["order"] = "true"
	data["n"] = "1000"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/playlist`, data, options)

	return reBody
}
