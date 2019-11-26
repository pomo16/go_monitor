package main

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/handler"
	"gowatcher/go_monitor/middleware"
)

var urlMaps = map[string]gin.HandlerFunc {
	"/ping": handler.Ping,
}

func InstanceRoutine() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(middleware.CustomLogger())
	for url, handler := range urlMaps {
		r.GET(url, handler)
		r.POST(url, handler)
	}
	return r
}
