package main

import (
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/service/redis"
)

func Init() {
	database.InitDB()
	redis.InitRedis()
}

func main() {
	Init()
	if err := InstanceRoutine().Run(":8080"); err != nil {
		panic(err)
	}
}
