package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type PlaylistCreateService struct {
	Name    string `json:"name" form:"name"`
	Privacy string `json:"privacy" form:"privacy"`
}

func (service *PlaylistCreateService) PlaylistCreate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Privacy != "10" {
		service.Privacy = "0"
	}
	data["name"] = service.Name
	data["privacy"] = service.Privacy
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/playlist/create`, data, options)

	return reBody
}
