package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjDetailService struct {
	ID string `json:"rid" form:"rid"`
}

func (service *DjDetailService) DjDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/get`, data, options)

	return reBody
}
