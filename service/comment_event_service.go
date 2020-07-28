package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type CommentEventService struct {
	ThreadId   string `json:"threadId" form:"threadId"`
	Limit      string `json:"limit" form:"limit"`
	Offset     string `json:"offset" form:"offset"`
	BeforeTime string `json:"beforeTime" form:"beforeTime"`
}

func (service *CommentEventService) CommentEvent(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "20"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	if service.BeforeTime == "" {
		data["beforeTime"] = "0"
	} else {
		data["beforeTime"] = service.BeforeTime
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/resource/comments/`+service.ThreadId, data, options)

	return reBody
}
