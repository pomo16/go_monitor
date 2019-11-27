package database

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"log"
)

func AddTask(ctx *gin.Context, params *model.CrawlParams) error {
	//todo: add sql
	return nil
}

func ListTask(ctx *gin.Context) ([]*model.TaskRow, error) {
	rows, err := database.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, app_id, status, create_time, modify_time").Order("id").Rows()

	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("Recovered in QueryTasks: %v\n", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		log.Printf("QueryTasks error, err: %v\n", err.Error())
		return nil, exceptions.ErrDBHandle
	}

	res := []*model.TaskRow{}
	for rows.Next() {
		var tmp model.TaskRow
		database.ScanRows(rows, &tmp)
		res = append(res, &tmp)
	}
	return res, nil
}

func GetTask(ctx *gin.Context, taskID int32) error {
	//todo: add sql
	return nil
}

func UpdateTask(ctx *gin.Context, params *model.CrawlParams) error {
	//todo: add sql
	return nil
}
