package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type FmTrashService struct {
	SongID string `json:"id" form:"id"`
}

func (service *FmTrashService) FmTrash(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)
	data["songId"] = service.SongID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/radio/trash/add?alg=RT&songId=`+service.SongID+`&time=25`, data, options)

	return reBody
}
