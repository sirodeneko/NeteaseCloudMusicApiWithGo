package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type LogoutService struct {
}

func (service *LogoutService) Logout(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/logout`, data, options)

	for _, cookie := range cookies {
		c.SetCookie(cookie.Name, cookie.Value, -1, "", cookie.Domain, false, false)
	}

	return reBody
}
