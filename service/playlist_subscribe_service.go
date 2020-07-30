package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistSubscribeService struct {
	T  string `json:"t" form:"t"`
	ID string `json:"id" form:"id"`
}

func (service *PlaylistSubscribeService) PlaylistSubscribe(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	if service.T == "1" {
		service.T = "subscribe"
	} else {
		service.T = "unsubscribe"
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/`+service.T, data, options)

	return reBody
}
