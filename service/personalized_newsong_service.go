package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PersonalizedNewsongService struct {
}

func (service *PersonalizedNewsongService) PersonalizedNewsong(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	data["type"] = "recommend"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/newsong`, data, options)

	return reBody
}
