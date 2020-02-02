package main

import (
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/redis"
)

func Init() {
	//database.InitDB()
	redis.InitRedis()
	elasticsearch.InitElasticSearch()
}

func main() {
	Init()
	if err := InstanceRoutine().Run(":8080"); err != nil {
		panic(err)
	}
}
