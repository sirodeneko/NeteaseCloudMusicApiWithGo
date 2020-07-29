package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ArtistTopSongService struct {
	Id string `json:"id" form:"id"`
}

func (service *ArtistTopSongService) ArtistTopSong(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	data["id"] = service.Id

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/artist/top/song`, data, options)

	return reBody
}
