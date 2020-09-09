package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjTodayPerferedService struct {
	Page string `json:"page" form:"page"`
}

func (service *DjTodayPerferedService) DjTodayPerfered(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Page == "" {
		data["page"] = "0"
	} else {
		data["page"] = service.Page
	}
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/home/today/perfered`, data, options)

	return reBody
}
