package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type RegisterCellphoneService struct {
	Phone    string `json:"phone" form:"phone"`
	Captcha  string `json:"captcha" form:"captcha"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
}

func (service *RegisterCellphoneService) RegisterCellphone(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	cookies = append(cookies, cookiesOS)

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: cookies,
	}
	data := make(map[string]string)

	data["phone"] = service.Phone
	h := md5.New()
	h.Write([]byte(service.Password))
	data["password"] = hex.EncodeToString(h.Sum(nil))
	data["captcha"] = service.Captcha
	data["nickname"] = service.Nickname

	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/register/cellphone`, data, options)

	return reBody
}
