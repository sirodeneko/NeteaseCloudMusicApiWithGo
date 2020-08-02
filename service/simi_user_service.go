package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SimiUserService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *SimiUserService) SimiUser(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["songid"] = service.ID
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/discovery/simiUser`, data, options)

	return reBody
}
