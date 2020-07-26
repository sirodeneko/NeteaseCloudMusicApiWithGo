package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

// LoginStatusService
type LoginStatusService struct{}

func (service *LoginStatusService) LoginStatus(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options:=&util.Options{
		Crypto:  "",
		Ua:      "",
		Cookies: cookies,
	}
	data:=make(map[string]string)
	reBody,cookies:=util.CreateRequest("GET",`https://music.163.com`,data,options)

	for _,cookie:=range cookies{
		c.SetCookie(cookie.Name,cookie.Value,60*60*24,"",cookie.Domain,false,false)
	}

	return reBody
}
