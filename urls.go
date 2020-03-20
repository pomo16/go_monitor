package main

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/handler"
	"gowatcher/go_monitor/handler/comment"
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

	"/comment/list":      comment.List,
	"/comment/count":     comment.Count,
	"/comment/histogram": comment.Histogram,

	"/retk":   user.Refresh,
	"/logout": user.Logout,
}

func InstanceRoutine() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())
	r.Use(middleware.Cors)

	authGroup := r.Group("/api/v1/auth")
	for url, handler := range outlookUrls {
		authGroup.POST(url, handler)
	}

	testGroup := r.Group("/api/v1/test")
	testGroup.Use(middleware.CheckLogin())
	for url, handler := range testUrls {
		testGroup.GET(url, handler)
		testGroup.POST(url, handler)
	}

	monitorGroup := r.Group("/api/v1/monitor")
	monitorGroup.Use(middleware.CheckLogin())
	for url, handler := range platformUrls {
		monitorGroup.GET(url, handler)
		monitorGroup.POST(url, handler)
	}

	return r
}
