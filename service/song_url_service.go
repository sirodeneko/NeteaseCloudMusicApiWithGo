package service

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type SongUrlService struct {
	ID string `json:"id" form:"id"`
	Br string `json:"br" form:"br"`
}

func (service *SongUrlService) SongUrl(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	var f = false

	for _, item := range cookies {
		if item.Name == "MUSIC_U" {
			f = true
			break
		}
	}

	if f {
		cookieNuid := &http.Cookie{Name: "_ntes_nuid", Value: hex.EncodeToString([]byte(util.RandStringRunes(16)))}
		cookies = append(cookies, cookieNuid)
	}

	options := &util.Options{
		Crypto:  "linuxapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["ids"] = "[" + service.ID + "]"
	if service.Br == "" {
		service.Br = "999000"
	}
	data["br"] = service.Br
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/song/enhance/player/url`, data, options)

	return reBody
}
