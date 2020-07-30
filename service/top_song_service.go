package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type TopSongService struct {
	AreaId string `json:"type" form:"type"`
}

func (service *TopSongService) TopSong(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.AreaId == "" {
		service.AreaId = "0"
	}
	data["areaId"] = service.AreaId
	data["total"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/discovery/new/songs`, data, options)

	return reBody
}
