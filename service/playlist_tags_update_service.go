package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type PlaylistTagsUpdateService struct {
	Id   string `json:"id" form:"id"`
	Tags string `json:"tags" form:"tags"`
}

func (service *PlaylistTagsUpdateService) PlaylistTagsUpdate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/playlist/tags/update",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["tags"] = service.Tags
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/tags/update`, data, options)

	return reBody
}
