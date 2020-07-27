package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/")
	{
		v1.POST("ping", api.Ping)

		v1.GET("login/status", api.LoginStatus)
		v1.GET("login/cellphone", api.LoginCellphone)
		v1.GET("login", api.LoginEmail)
		v1.GET("login/refresh", api.LoginRefresh)
		v1.GET("captcha/sent", api.CaptchaSent)
		v1.GET("captcha/verify", api.CaptchaVerify)
		v1.GET("register/cellphone", api.RegisterCellphone)
		v1.GET("cellphone/existence/check", api.CellphoneExistenceCheck)
		v1.GET("activate/init/profile", api.ActivateInitProfile)
		v1.GET("rebind", api.Rebind)
		v1.GET("logout", api.Logout)
		v1.GET("user/detail", api.UserDetail)
		v1.GET("user/subcount", api.UserSubcount)
		v1.GET("user/update", api.UserUpdate)
		v1.GET("countries/code/list", api.CountriesCodeList)
		v1.GET("user/playlist", api.UserPlaylist)
		v1.GET("user/playlist/update", api.PlaylistUpdate)
		v1.GET("playlist/desc/update", api.PlaylistDescUpdate)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			//auth.GET("user/me", api.UserMe)
			//auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
