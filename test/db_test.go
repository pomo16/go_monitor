package test

import (
	"context"
	"fmt"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/utils"
	"testing"
)

func TestDB(t *testing.T) {
	database.InitDB()
	res, err := database.GetTaskList(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}

func TestInsertTask(t *testing.T) {
	database.InitDB()
	task := &model.CrawlTask{
		AppID:      "458318329",
		AppName:    "腾讯视频",
		Status:     1,
		CreateTime: "2019-11-26 10:50:55",
		ModifyTime: "2019-11-26 10:50:55",
	}
	err := database.InsertTask(context.Background(), task)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetTaskByID(t *testing.T) {
	database.InitDB()
	res, err := database.GetTaskByID(context.Background(), 8)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}

func TestUpdateTask(t *testing.T) {
	database.InitDB()
	task := &model.CrawlTask{
		ID:     7,
		Status: 2,
	}
	err := database.UpdateTask(context.Background(), task)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCheckUser(t *testing.T) {
	database.InitDB()
	params := &model.LoginParams{
		UserName: "pomo",
		Password: utils.Md5AddSalt("123", consts.PasswordSalt, false),
	}
	user, err := database.CheckUser(context.Background(), params)
	if err != nil {
		fmt.Errorf("err: %v", err)
		return
	}

	fmt.Printf("user: %v", user)
}
