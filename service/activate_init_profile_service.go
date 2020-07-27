package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ActivateInitProfileService struct {
	Nickname string `json:"nickname" form:"nickname"`
}

func (service *ActivateInitProfileService) ActivateInitProfile(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/activate/initProfile",
	}
	data := make(map[string]string)
	data["nickname"] = service.Nickname

	reBody, _ := util.CreateRequest("POST", `http://music.163.com/eapi/activate/initProfile`, data, options)

	return reBody
}
