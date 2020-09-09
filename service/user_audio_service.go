package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserAudioService struct {
	UID string `json:"uid" form:"uid"`
}

func (service *UserAudioService) UserAudio(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["userId"] = service.UID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/djradio/get/byuser`, data, options)

	return reBody
}
