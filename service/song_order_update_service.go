package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type SongOrderUpdateService struct {
	Pid string `json:"pid" form:"pid"`
	Ids string `json:"ids" form:"ids"`
}

func (service *SongOrderUpdateService) SongOrderUpdate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
		Url:     "/api/playlist/desc/update",
	}
	data := make(map[string]string)
	data["pid"] = service.Pid
	data["trackIds"] = service.Ids
	data["op"] = "update"
	reBody, _ := util.CreateRequest("POST", `http://interface.music.163.com/api/playlist/manipulate/tracks`, data, options)

	return reBody
}
