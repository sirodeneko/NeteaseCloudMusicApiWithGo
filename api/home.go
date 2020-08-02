package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// 首页-发现 block page
// 这个接口为移动端接口，首页-发现页，数据结构可以参考 https://github.com/hcanyz/flutter-netease-music-api/blob/master/lib/src/api/uncategorized/bean.dart#L259 HomeBlockPageWrap
// query.refresh 是否刷新数据
func HomepageBlockPage(c *gin.Context) {
	var service service.HomepageBlockPageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HomepageBlockPage(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 首页-发现 dragon ball
// 这个接口为移动端接口，首页-发现页（每日推荐、歌单、排行榜 那些入口）
// 数据结构可以参考 https://github.com/hcanyz/flutter-netease-music-api/blob/master/lib/src/api/uncategorized/bean.dart#L290 HomeDragonBallWrap
// !需要登录或者匿名登录，非登录返回 []
func HomepageDragonBall(c *gin.Context) {
	var service service.HomepageDragonBallService
	if err := c.ShouldBind(&service); err == nil {
		res := service.HomepageDragonBall(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
