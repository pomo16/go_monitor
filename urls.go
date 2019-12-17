package main

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/handler"
	"gowatcher/go_monitor/handler/crawl"
	"gowatcher/go_monitor/handler/user"
	"gowatcher/go_monitor/middleware"
)

//对外暴露url，不强制校验登录态
var outlookUrls = map[string]gin.HandlerFunc{
	"/login": user.Login,
}

var testUrls = map[string]gin.HandlerFunc{
	"/ping": handler.Ping,
}

var platformUrls = map[string]gin.HandlerFunc{
	"/crawl/config": crawl.TaskConfig,
	"/crawl/list":   crawl.TaskList,
	"/logout":       user.Logout,
}

func InstanceRoutine() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	authGroup := r.Group("/auth")
	for url, handler := range outlookUrls {
		authGroup.GET(url, handler)
		authGroup.POST(url, handler)
	}

	testGroup := r.Group("/test")
	testGroup.Use(middleware.CheckLogin())
	for url, handler := range testUrls {
		testGroup.GET(url, handler)
		testGroup.POST(url, handler)
	}

	monitorGroup := r.Group("/monitor")
	for url, handler := range platformUrls {
		monitorGroup.GET(url, handler)
		monitorGroup.POST(url, handler)
	}

	return r
}
