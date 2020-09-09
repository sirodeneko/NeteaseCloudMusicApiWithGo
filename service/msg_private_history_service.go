package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type MsgPrivateHistoryService struct {
	UID   string `json:"uid" form:"uid"`
	Limit string `json:"limit" form:"limit"`
	Time  string `json:"before" form:"before"`
}

func (service *MsgPrivateHistoryService) MsgPrivateHistory(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["userId"] = service.UID
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Time == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Time
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/msg/private/history`, data, options)

	return reBody
}
