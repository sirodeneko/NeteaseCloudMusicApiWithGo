package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type ArtistDescService struct {
	ID string `json:"id" form:"id"`
}

func (service *ArtistDescService) ArtistDesc(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/introduction`, data, options)

	return reBody
}
