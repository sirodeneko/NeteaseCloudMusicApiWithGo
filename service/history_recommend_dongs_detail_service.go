package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type HistoryRecommendDongsDetailService struct {
	Date string `json:"date" form:"date"`
}

func (service *HistoryRecommendDongsDetailService) HistoryRecommendDongsDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "ios"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["date"] = service.Date

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/discovery/recommend/songs/history/detail`, data, options)

	return reBody
}
