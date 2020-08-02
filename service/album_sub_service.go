package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type AlbumSubService struct {
	ID string `json:"id" form:"id"`
	T  string `json:"t" form:"t"`
}

func (service *AlbumSubService) AlbumSub(c *gin.Context) map[string]interface{} {

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
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/album/`+service.T, data, options)

	return reBody
}
