package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type CheckMusicService struct {
	ID string `json:"id" form:"id"`
	Br string `json:"br" form:"br"`
}

func (service *CheckMusicService) CheckMusic(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["ids"] = "[" + service.ID + "]"
	if service.Br == "" {
		service.Br = "999000"
	}
	data["br"] = service.Br
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/song/enhance/player/url`, data, options)

	if reBody["code"].(float64) == 200 {
		reBody["success"] = true
		reBody["message"] = "ok"
	} else {
		reBody["success"] = false
		reBody["message"] = "亲爱的,暂无版权"
	}
	return reBody
}
