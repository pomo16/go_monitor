package main

import (
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/redis"
)

func Init() {
	database.InitDB()
	redis.InitRedis()
	elasticsearch.InitElasticSearch()
}

func main() {
	Init()
	if err := InstanceRoutine().Run("112.74.86.176:8888"); err != nil {
		panic(err)
	}
}
