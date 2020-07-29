package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistHotService struct{}

func (service *PlaylistHotService) PlaylistHot(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/hottags`, data, options)

	return reBody
}
