package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type RelatedAllVideoService struct {
	ID string `json:"id" form:"id"`
}

func (service *RelatedAllVideoService) RelatedAllVideo(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["type"] = "1"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/cloudvideo/v1/allvideo/rcmd`, data, options)

	return reBody
}
