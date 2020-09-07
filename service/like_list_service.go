package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type LikeListService struct {
	UID string `json:"uid" form:"uid"`
}

func (service *LikeListService) LikeList(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}

	data := make(map[string]string)
	data["uid"] = service.UID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/song/like/get`, data, options)

	return reBody
}
