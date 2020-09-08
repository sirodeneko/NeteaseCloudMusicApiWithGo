package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type VideoTimelineRecommendService struct {
	Offset string `json:"offset" form:"offset"`
}

func (service *VideoTimelineRecommendService) VideoTimelineRecommend(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["filterLives"] = "[]"
	data["withProgramInfo"] = "true"
	data["needUrl"] = "1"
	data["resolution"] = "480"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/videotimeline/get`, data, options)

	return reBody
}
