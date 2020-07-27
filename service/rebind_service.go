package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type RebindService struct {
	Captcha    string `json:"captcha" form:"captcha"`
	Phone      string `json:"phone" form:"phone"`
	Oldcaptcha string `json:"oldcaptcha" form:"oldcaptcha"`
	Ctcode     string `json:"ctcode" form:"ctcode"`
}

func (service *RebindService) Rebind(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	data["phone"] = service.Phone
	data["captcha"] = service.Captcha
	data["captcha"] = service.Captcha
	data["oldcaptcha"] = service.Oldcaptcha

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/user/replaceCellphone`, data, options)

	return reBody
}
