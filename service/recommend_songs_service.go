package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type RecommendSongsService struct {
	ID string `json:"id" form:"id"`
}

func (service *RecommendSongsService) RecommendSongs(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "ios"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/v3/discovery/recommend/songs`, data, options)

	return reBody
}
