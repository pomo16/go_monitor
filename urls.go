package main

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/handler"
	"gowatcher/go_monitor/handler/crawl"
	"gowatcher/go_monitor/middleware"
)

var testUrls = map[string]gin.HandlerFunc{
	"/ping": handler.Ping,
}

var platformUrls = map[string]gin.HandlerFunc{
	"/crawl/add":    crawl.AddTask,
	"/crawl/list":   crawl.TaskList,
	"/crawl/get":    crawl.GetTask,
	"/crawl/update": crawl.UpdateTask,
}

func InstanceRoutine() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(middleware.CustomLogger())

	testGroup := r.Group("/test")
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
