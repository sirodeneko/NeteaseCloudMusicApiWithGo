package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjToplistPayService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *DjToplistPayService) DjToplistPay(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Limit == "" {
		data["limit"] = "10"
	} else {
		data["limit"] = service.Limit
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/djradio/toplist/pay`, data, options)

	return reBody
}
