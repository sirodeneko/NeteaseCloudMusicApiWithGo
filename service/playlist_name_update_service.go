package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistNameUpdateService struct {
	Id   string `json:"id" form:"id"`
	Name string `json:"desc" form:"name"`
}

func (service *PlaylistNameUpdateService) PlaylistNameUpdate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/playlist/update/name",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["name"] = service.Name
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/update/name`, data, options)

	return reBody
}
