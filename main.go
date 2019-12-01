package main

import "gowatcher/go_monitor/service/database"

func Init() {
	database.InitDB()
}

func main() {
	Init()
	if err := InstanceRoutine().Run(":8080"); err != nil {
		panic(err)
	}
}
