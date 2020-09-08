package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type VideoTimelineAllService struct {
	Offset string `json:"offset" form:"offset"`
}

func (service *VideoTimelineAllService) VideoTimelineAll(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["groupId"] = "0"
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	data["need_preview_url"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/videotimeline/otherclient/get`, data, options)

	return reBody
}
