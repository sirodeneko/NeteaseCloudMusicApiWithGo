package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type LoginRefreshService struct {
}

func (service *LoginRefreshService) LoginRefresh(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/login/token/refresh`, data, options)

	return reBody
}
