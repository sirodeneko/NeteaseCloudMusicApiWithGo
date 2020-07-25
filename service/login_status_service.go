package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// LoginStatusService
type LoginStatusService struct{}

func (service *LoginStatusService) LoginStatus(c *gin.Context) serializer.Response {

	// 获得所有cookie
	cookies := c.Request.Cookies()
	c

	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
