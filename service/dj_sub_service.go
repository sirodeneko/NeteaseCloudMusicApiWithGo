package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjSubService struct {
	RID string `json:"rid" form:"rid"`
	T   string `json:"t" form:"t"`
}

func (service *DjSubService) DjSub(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.RID
	if service.T == "1" {
		service.T = "sub"
	} else {
		service.T = "unsub"
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/`+service.T, data, options)

	return reBody
}
