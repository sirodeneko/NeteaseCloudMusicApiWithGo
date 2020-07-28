package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type PlaylistOrderUpdateService struct {
	Ids string `json:"ids" form:"ids"`
}

func (service *PlaylistOrderUpdateService) PlaylistOrderUpdate(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["id"] = service.Ids
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/playlist/order/update`, data, options)

	return reBody
}
