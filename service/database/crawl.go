package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//InsertTask 添加爬虫任务
func InsertTask(ctx context.Context, task *model.CrawlTask) error {
	if err := database.Table("gowatcher.crawl_task_table").Debug().Create(task).Error; err != nil {
		logrus.Warnf("insert task error, err: %v", err.Error())
		return exceptions.ErrDBHandle
	}
	return nil
}

//GetTaskList 获取爬虫任务列表
func GetTaskList(ctx context.Context) ([]*model.CrawlTask, error) {
	rows, err := database.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, app_id, status, create_time, modify_time").Order("id").Rows()

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
		var tmp model.CrawlTask
		database.ScanRows(rows, &tmp)
		res = append(res, &tmp)
	}
	return res, nil
}

//GetTaskByID 通过ID获取爬虫任务详情
func GetTaskByID(ctx context.Context, taskID int32) ([]*model.CrawlTask, error) {
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
		var tmp model.CrawlTask
		database.ScanRows(rows, &tmp)
		res = append(res, &tmp)
	}
	return res, nil
}

//UpdateTask 更新爬虫任务状态
func UpdateTask(ctx context.Context, task *model.CrawlTask) error {
	err := database.Table("gowatcher.crawl_task_table").Debug().Where("id = ?", task.ID).
		UpdateColumns(model.CrawlTask{Status: task.Status, ModifyTime: task.ModifyTime}).Error

	if err != nil {
		logrus.Errorf("update task error, err: %v", err.Error())
		return exceptions.ErrDBHandle
	}
	return nil
}
