package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ProgramRecommendService struct {
	CateId string `json:"type" form:"type"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *ProgramRecommendService) ProgramRecommend(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["cateId"] = service.CateId
	if service.Limit == "" {
		data["limit"] = "10"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/program/recommend/v1`, data, options)

	return reBody
}
