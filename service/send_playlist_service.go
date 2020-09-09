package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type SendPlaylistService struct {
	ID      string `json:"playlist" form:"playlist"`
	Msg     string `json:"msg" form:"msg"`
	UserIds string `json:"user_ids" form:"user_ids"`
}

func (service *SendPlaylistService) SendPlaylist(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["type"] = "playlist"
	data["msg"] = service.Msg
	data["userIds"] = "[" + service.UserIds + "]"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/msg/private/send`, data, options)

	return reBody
}
