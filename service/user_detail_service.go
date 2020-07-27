package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserDetailService struct {
	Uid string `json:"uid" form:"uid"`
}

func (service *UserDetailService) UserDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/v1/user/detail/`+service.Uid, data, options)

	return reBody
}
