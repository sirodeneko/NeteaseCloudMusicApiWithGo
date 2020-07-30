package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistTracksService struct {
	Op       string `json:"op" form:"op"`
	Pid      string `json:"pid" form:"pid"`
	TrackIds string `json:"tracks" form:"tracks"`
}

func (service *PlaylistTracksService) PlaylistTracks(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["op"] = service.Op
	data["pid"] = service.Pid
	data["trackIds"] = "[" + service.TrackIds + "]"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/manipulate/tracks`, data, options)

	return reBody
}
