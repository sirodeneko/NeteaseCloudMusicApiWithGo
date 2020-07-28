package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type EventDelService struct {
	EvId string `json:"evId" form:"evId"`
}

func (service *EventDelService) EventDel(c *gin.Context) map[string]interface{} {

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

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/eapi/event/delete`, data, options)

	return reBody
}
