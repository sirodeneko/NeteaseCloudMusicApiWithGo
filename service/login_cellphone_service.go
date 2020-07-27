package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/util"
)

type LoginCellphoneService struct {
	Phone       string `json:"phone" form:"phone"`
	Countrycode string `json:"countrycode" form:"countrycode"`
	Password    string `json:"password" form:"password"`
	Md5password string `json:"md5_password" form:"md5_password"`
}

func (service *LoginCellphoneService) LoginCellphone(c *gin.Context) map[string]interface{} {

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

	data["phone"] = service.Phone
	if service.Countrycode != "" {
		data["countrycode"] = service.Countrycode
	}
	if service.Password != "" {
		h := md5.New()
		h.Write([]byte(service.Password))
		data["password"] = hex.EncodeToString(h.Sum(nil))
	} else {
		data["password"] = service.Md5password
	}
	data["rememberLogin"] = "true"

	//reBody, cookies := util.CreateRequest("POST", `https://www.httpbin.org/post`, data, options)
	reBody, cookies := util.CreateRequest("POST", `https://music.163.com/weapi/login/cellphone`, data, options)

	cookiesStr := ""

	for _, cookie := range cookies {
		if cookiesStr != "" {
			cookiesStr = cookiesStr + ";"
		}
		cookiesStr = cookiesStr + cookie.String()
		c.SetCookie(cookie.Name, cookie.Value, 60*60*24, "", cookie.Domain, false, false)
	}

	reBody["cookie"] = cookiesStr

	return reBody
}
