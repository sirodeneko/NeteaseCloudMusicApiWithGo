package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistDetailService struct {
	Id string `json:"id" form:"id"`
	S  string `json:"s" form:"s"`
}

func (service *PlaylistDetailService) PlaylistDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "linuxapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.S == "" {
		service.S = "8"
	}
	data["id"] = service.Id
	data["n"] = "100000"
	data["s"] = service.S

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v3/playlist/detail`, data, options)

	return reBody
}
