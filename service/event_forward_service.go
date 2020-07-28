package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type EventForwardService struct {
	Uid      string `json:"uid" form:"uid"`
	EvId     string `json:"evId" form:"evId"`
	Forwards string `json:"forwards" form:"forwards"`
}

func (service *EventForwardService) EventForward(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.EvId
	data["eventUserId"] = service.Uid
	data["forwards"] = service.Forwards
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/event/forward`, data, options)

	return reBody
}
