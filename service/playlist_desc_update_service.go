package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistDescUpdateService struct {
	Id   string `json:"id" form:"id"`
	Desc string `json:"desc" form:"desc"`
}

func (service *PlaylistDescUpdateService) PlaylistDescUpdate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/playlist/desc/update",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["desc"] = service.Desc
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/desc/update`, data, options)

	return reBody
}
