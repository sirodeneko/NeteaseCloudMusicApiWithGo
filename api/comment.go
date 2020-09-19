package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CommentEvent(c *gin.Context) {
	var service service.CommentEventService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentEvent(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 云村热评
func CommentHotwallList(c *gin.Context) {
	var service service.CommentHotwallListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentHotwallList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 音乐评论
func CommentMusic(c *gin.Context) {
	var service service.CommentMusicService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentMusic(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 专辑评论
func CommentAlbum(c *gin.Context) {
	var service service.CommentAlbumService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentAlbum(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌单评论
func CommentPlaylist(c *gin.Context) {
	var service service.CommentPlaylistService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentPlaylist(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// mv评论
func CommentMv(c *gin.Context) {
	var service service.CommentMvService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentMv(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 电台评论
func CommentDj(c *gin.Context) {
	var service service.CommentDjService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentDj(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频评论
func CommentVideo(c *gin.Context) {
	var service service.CommentVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentVideo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 热门评论
func CommentHot(c *gin.Context) {
	var service service.CommentHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentHot(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 评论点赞
func CommentLike(c *gin.Context) {
	var service service.CommentLikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentLike(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 评论
func Comment(c *gin.Context) {
	var service service.CommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Comment(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 歌曲楼层评论
func CommentFloor(c *gin.Context) {
	var service service.CommentFloorService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommentFloor(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
