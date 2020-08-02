package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type LikeService struct {
	ID string `json:"id" form:"id"`
	L  string `json:"like" form:"like"`
}

func (service *LikeService) Like(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["trackId"] = service.ID
	if service.L == "" {
		data["like"] = "true"
	} else {
		data["like"] = service.L
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/radio/like?alg=itembased&trackId=`+service.ID+`&time=25`, data, options)

	return reBody
}
