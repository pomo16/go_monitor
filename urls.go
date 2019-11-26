package main

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/middleware"
	"lhx-github/go_web_demo/handler"
)

var urlMaps = map[string]gin.HandlerFunc {
	"/ping": handler.Ping,
}

func InstanceRoutine() *gin.Engine {
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()
	r.Use(middleware.CustomLogger())
	for url, handler := range urlMaps {
		r.GET(url, handler)
		r.POST(url, handler)
	}
	return r
}
