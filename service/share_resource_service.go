package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ShareResourceService struct {
	Id   string `json:"id" form:"id"`
	Msg  string `json:"msg" form:"msg"`
	Type string `json:"type" form:"type"`
}

func (service *ShareResourceService) ShareResource(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["msg"] = service.Msg

	if service.Type == "" {
		data["type"] = "song"
	} else {
		data["type"] = service.Type
	}
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/share/friends/resource`, data, options)

	return reBody
}
