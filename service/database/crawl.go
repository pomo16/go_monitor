package database

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

func AddTask(ctx *gin.Context, params *model.CrawlParams) error {
	//todo: add sql
	return nil
}

func GetTaskList(ctx *gin.Context) ([]*model.CrawlTask, error) {
	rows, err := database.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, status").Order("id").Rows()

	defer func() {
		if r := recover(); r != nil {
			logrus.Warn("Recovered in QueryTasks: %v\n", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		logrus.Errorf("QueryTasks error, err: %v\n", err.Error())
		return nil, exceptions.ErrDBHandle
	}

	res := []*model.CrawlTask{}
	for rows.Next() {
		var tmp model.TaskRow
		database.ScanRows(rows, &tmp)
		task := &model.CrawlTask{
			TaskID:     tmp.ID,
			AppID:      tmp.AppID,
			AppName:    tmp.AppName,
			Status:     tmp.Status,
			CreateTime: tmp.CreateTime,
			ModifyTime: tmp.ModifyTime,
		}
		res = append(res, task)
	}
	return res, nil
}

func GetTaskByID(ctx *gin.Context, taskID int32) ([]*model.CrawlTask, error) {
	rows, err := database.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, app_id, status, create_time, modify_time").
		Where("id = ?", taskID).Rows()

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in QueryTasks: %v\n", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		logrus.Errorf("QueryTasks error, err: %v\n", err.Error())
		return nil, exceptions.ErrDBHandle
	}

	res := []*model.CrawlTask{}
	for rows.Next() {
		var tmp model.TaskRow
		database.ScanRows(rows, &tmp)
		task := &model.CrawlTask{
			TaskID:     tmp.ID,
			AppID:      tmp.AppID,
			AppName:    tmp.AppName,
			Status:     tmp.Status,
			CreateTime: tmp.CreateTime,
			ModifyTime: tmp.ModifyTime,
		}
		res = append(res, task)
	}
	return res, nil
}

func UpdateTask(ctx *gin.Context, params *model.CrawlParams) error {
	//todo: add sql
	return nil
}
