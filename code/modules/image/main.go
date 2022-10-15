package main

import (
	"duanxu.tech/website/bilibili"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	image := router.Group("/image")
	{
		image.GET("/list", bilibili.GetVideoList)
	}

	router.Run(":9000")
}
