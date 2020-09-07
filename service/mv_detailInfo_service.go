package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type MvDetailInfoService struct {
	ID string `json:"mvid" form:"mvid"`
}

func (service *MvDetailInfoService) MvDetailInfo(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["threadid"] = "R_MV_5_" + service.ID
	data["composeliked"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/comment/commentthread/info`, data, options)

	return reBody
}
