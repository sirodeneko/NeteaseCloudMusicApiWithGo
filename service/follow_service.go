package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type FollowService struct {
	T  string `json:"t" form:"t"`
	Id string `json:"id" form:"id"`
}

func (service *FollowService) Follow(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	if service.T == "1" {
		service.T = "follow"
	} else {
		service.T = "delfollow"
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/user/`+service.T+`/`+service.Id, data, options)

	return reBody
}
