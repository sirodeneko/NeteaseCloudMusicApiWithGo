package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type VideoUrlService struct {
	ID  string `json:"id" form:"id"`
	Res string `json:"resolution" form:"resolution"`
}

func (service *VideoUrlService) VideoUrl(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Res == "" {
		data["resolution"] = "1080"
	} else {
		data["resolution"] = service.Res
	}
	data["ids"] = `["` + service.ID + `"]`
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/cloudvideo/playurl`, data, options)

	return reBody
}
