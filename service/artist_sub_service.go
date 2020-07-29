package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ArtistSubService struct {
	T  string `json:"t" form:"t"`
	Id string `json:"id" form:"id"`
}

func (service *ArtistSubService) ArtistSub(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	if service.T == "1" {
		service.T = "sub"
	} else {
		service.T = "unsub"
	}

	data["artistId"] = service.Id
	data["artistIds"] = "[" + service.Id + "]"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/`+service.T, data, options)

	return reBody
}
