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
	}

	router.Run(":9000")
}
