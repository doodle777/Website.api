package main

import (
	"duanxu.tech/website/bilibili"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	video := router.Group("/video")
	{
		video.GET("/list", bilibili.GetVideoList)
		video.GET("/detail", bilibili.GetVideoDetail)
	}

	_ = router.Run(":9000")
}
