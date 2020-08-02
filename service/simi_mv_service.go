package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SimiMvService struct {
	ID     string `json:"mvid" form:"mvid"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *SimiMvService) SimiMv(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["mvid"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/discovery/simiMV`, data, options)

	return reBody
}
