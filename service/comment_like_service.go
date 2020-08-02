package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type CommentLikeService struct {
	ID       string `json:"id" form:"id"`
	ThreadId string `json:"threadId" form:"threadId"`
	Cid      string `json:"cid" form:"cid"`
	T        string `json:"t" form:"t"`
	Type     string `json:"type" form:"type"`
}

func (service *CommentLikeService) CommentLike(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}

	TYPE := make(map[string]string, 6)
	TYPE["0"] = "R_SO_4_"
	TYPE["1"] = "R_MV_5_"
	TYPE["2"] = "A_PL_0_"
	TYPE["3"] = "R_AL_3_"
	TYPE["4"] = "A_DJ_1_"
	TYPE["5"] = "R_VI_62_"
	TYPE["6"] = "A_EV_2_"

	if _, ok := TYPE[service.Type]; ok {
		service.Type = TYPE[service.Type]
	} else {
		service.Type = TYPE["0"]
	}

	data := make(map[string]string)
	data["commentId"] = service.Cid
	data["threadId"] = service.Type + service.ID

	if service.Type == "A_EV_2_" {
		data["threadId"] = service.ThreadId
	}

	if service.T == "1" {
		service.T = "like"
	} else {
		service.T = "unlike"
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/comment/`+service.T, data, options)

	return reBody
}
