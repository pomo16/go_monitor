package crawl

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
)

//TaskSetter 爬虫任务更新器
type TaskSetter struct{}

//Process 更新爬虫任务
func (setter *TaskSetter) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	inputParameter := runCtx.GetInputParameter()
	if inputParameter.Status < consts.Normal || inputParameter.Status > consts.Unused {
		return exceptions.ErrRequestParams
	}

	params := &model.CrawlParams{
		AppID:   inputParameter.AppID,
		AppName: inputParameter.AppName,
		Status:  inputParameter.Status,
	}

	if err := database.UpdateTask(ctx, params); err == nil {
		return nil
	}

	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (setter *TaskSetter) Name() string {
	return "TaskSetter"
}
