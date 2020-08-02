package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DailySigninService struct {
	Type string `json:"type" form:"type"`
}

func (service *DailySigninService) DailySignin(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.Type == "" {
		data["type"] = "0"
	} else {
		data["type"] = service.Type
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/point/dailyTask`, data, options)

	return reBody
}
