package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type PlaylistDeleteService struct {
	ID string `json:"id" form:"id"`
}

func (service *PlaylistDeleteService) PlaylistDelete(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["ids"] = "[" + service.ID + "]"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/remove`, data, options)

	return reBody
}
