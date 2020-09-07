package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type MvDetailService struct {
	ID string `json:"mvid" form:"mvid"`
}

func (service *MvDetailService) MvDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v1/mv/detail`, data, options)

	return reBody
}
