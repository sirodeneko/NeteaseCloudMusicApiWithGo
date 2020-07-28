package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserRecordService struct {
	UId  string `json:"uid" form:"uid"`
	Type string `json:"type" form:"type"`
}

func (service *UserRecordService) UserRecord(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["uid"] = service.UId

	if service.Type == "1" {
		data["type"] = "1"
	} else {
		data["type"] = "0"
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/play/record`, data, options)

	return reBody
}
