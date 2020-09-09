package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserCloudDetailService struct {
	ID string `json:"id" form:"id"`
}

func (service *UserCloudDetailService) UserCloudDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["songIds"] = "[" + service.ID + "]"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/cloud/get/byids`, data, options)

	return reBody
}
