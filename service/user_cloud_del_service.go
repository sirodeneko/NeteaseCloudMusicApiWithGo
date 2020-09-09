package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserCloudDelService struct {
	ID string `json:"id" form:"id"`
}

func (service *UserCloudDelService) UserCloudDel(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["songIds"] = "[" + service.ID + "]"
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/cloud/del`, data, options)

	return reBody
}
