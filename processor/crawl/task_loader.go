package crawl

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
)

//TaskLoader 爬虫任务加载器
type TaskLoader struct{}

//Process 加载爬虫任务
func (loader *TaskLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	listCtx, ok := runCtx.(model.ITaskListContext)
	if !ok {
		logrus.Warn(ctx, "TaskList loader listCtx error")
		return exceptions.ErrTypeAssert
	}

	inputParameter := runCtx.GetInputParameter()
	var taskList []*model.CrawlTask
	var err error
	switch inputParameter.QueryType {
	case consts.IdType:
		if inputParameter.TaskID == 0 {
			return exceptions.ErrRequestParams
		} else {
			taskList, err = database.GetTaskByID(ctx, inputParameter.TaskID)
		}
	case consts.ListType:
		taskList, err = database.GetTaskList(ctx)
	default:
		return exceptions.ErrRequestParams
	}

	if taskList != nil {
		listCtx.SetTaskList(taskList)
		return nil
	}

	if err == nil && len(taskList) == 0 {
		return exceptions.ErrResultEmpty
	}

	logrus.Error("TaskList loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *TaskLoader) Name() string {
	return "TaskLoader"
}
