package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjCategoryExcludehotService struct {
}

func (service *DjCategoryExcludehotService) DjCategoryExcludehot(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/category/excludehot`, data, options)

	return reBody
}
