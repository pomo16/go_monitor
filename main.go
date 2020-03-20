package main

import (
	"gowatcher/go_monitor/algoml"
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/redis"
)

func Init() {
	database.InitDB()
	redis.InitRedis()
	algoml.InitAlgoModel()
	elasticsearch.InitElasticSearch()
}

func main() {
	Init()
	if err := InstanceRoutine().Run(":8888"); err != nil {
		panic(err)
	}
}
