package service

import (
	"github.com/gin-gonic/gin"
	"singo/util"
)

type CellphoneExistenceCheckService struct {
	Cellphone   string `json:"phone" form:"phone"`
	Countrycode string `json:"countrycode" form:"countrycode"`
}

func (service *CellphoneExistenceCheckService) CellphoneExistenceCheck(c *gin.Context) map[string]interface{} {

	// 获得所有cookie
	cookies := c.Request.Cookies()

	options := &util.Options{
		Crypto:  "eapi",
		Cookies: cookies,
		Url:     "/api/cellphone/existence/check",
	}
	data := make(map[string]string)
	if service.Countrycode != "" {
		data["countrycode"] = service.Countrycode
	}
	data["cellphone"] = service.Cellphone

	reBody, cookies := util.CreateRequest("POST", `http://music.163.com/eapi/cellphone/existence/check`, data, options)

	return reBody
}
