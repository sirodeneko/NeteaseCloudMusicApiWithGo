package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type LyricService struct {
	ID string `json:"id" form:"id"`
}

func (service *LyricService) Lyric(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "linuxapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["lv"] = "-1"
	data["kv"] = "-1"
	data["tv"] = "-1"

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/song/lyric`, data, options)

	return reBody
}
