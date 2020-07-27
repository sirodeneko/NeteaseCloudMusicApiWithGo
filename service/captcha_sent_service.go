package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type CaptchaSentService struct {
	Ctcode    string `json:"ctcode" form:"ctcode"`
	Cellphone string `json:"phone" form:"phone"`
}

func (service *CaptchaSentService) CaptchaSent(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Ctcode == "" {
		data["ctcode"] = "86"
	} else {
		data["ctcode"] = service.Ctcode
	}
	data["cellphone"] = service.Cellphone

	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/sms/captcha/sent`, data, options)

	return reBody
}
