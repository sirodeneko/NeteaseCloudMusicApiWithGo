package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type EventService struct {
	PageSize string `json:"pagesize" form:"pagesize"`
	LastTime string `json:"lasttime" form:"lasttime"`
}

func (service *EventService) Event(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.PageSize == "" {
		data["pagesize"] = "20"
	} else {
		data["pagesize"] = service.PageSize
	}
	if service.LastTime == "" {
		data["lasttime"] = "-1"
	} else {
		data["lasttime"] = service.LastTime
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/event/get`, data, options)

	return reBody
}
