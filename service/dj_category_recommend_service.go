package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type DjCategoryRecommendService struct {
}

func (service *DjCategoryRecommendService) DjCategoryRecommend(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	reBody, _ := util.CreateRequest("POST", `http://music.163.com/weapi/djradio/home/category/recommend`, data, options)

	return reBody
}
