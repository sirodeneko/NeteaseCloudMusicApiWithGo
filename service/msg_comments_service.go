package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type MsgCommentsService struct {
	UID        string `json:"uid" form:"uid"`
	Limit      string `json:"limit" form:"limit"`
	BeforeTime string `json:"before" form:"before"`
}

func (service *MsgCommentsService) MsgComments(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["uid"] = service.UID
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.BeforeTime == "" {
		data["beforeTime"] = "-1"
	} else {
		data["beforeTime"] = service.BeforeTime
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v1/user/comments/`+service.UID, data, options)

	return reBody
}
