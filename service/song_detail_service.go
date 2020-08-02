package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
	"strings"
)

type SongDetailService struct {
	Ids string `json:"ids" form:"ids"`
}

func (service *SongDetailService) SongDetail(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}

	type IDS struct {
		ID string `json:"id"`
	}

	var cids []IDS

	strs := strings.Split(service.Ids, ",")
	for _, item := range strs {
		cids = append(cids, IDS{ID: item})
	}
	sidsJsonByte, _ := json.Marshal(cids)

	data := make(map[string]string)
	data["c"] = string(sidsJsonByte)
	data["ids"] = "[" + service.Ids + "]"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v3/song/detail`, data, options)

	return reBody
}
