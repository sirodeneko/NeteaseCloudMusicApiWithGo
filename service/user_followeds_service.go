package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type UserFollowedsService struct {
	Uid   string `json:"uid" form:"uid"`
	Limit string `json:"limit" form:"limit"`
	Time  string `json:"lasttime " form:"lasttime "`
}

func (service *UserFollowedsService) UserFolloweds(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/user/getfolloweds",
	}
	data := make(map[string]string)
	data["userId"] = service.Uid
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Time == "" {
		data["time"] = "-1"
	} else {
		data["time"] = service.Time
	}
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/eapi/user/getfolloweds/`+service.Uid, data, options)

	return reBody
}
