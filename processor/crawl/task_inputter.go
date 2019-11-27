package crawl

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
)

//TaskInputter 爬虫任务添加器
type TaskInputter struct{}

//Process 添加爬虫任务
func (inputter *TaskInputter) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	inputParameter := runCtx.GetInputParameter()
	if inputParameter.AppID == "" || inputParameter.AppName == "" || inputParameter.Status < consts.Normal || inputParameter.Status > consts.Unused {
		return exceptions.ErrRequestParams
	}

	params := &model.CrawlParams{
		AppID:   inputParameter.AppID,
		AppName: inputParameter.AppName,
		Status:  inputParameter.Status,
	}

	if err := database.AddTask(ctx, params); err == nil {
		return nil
	}

	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (inputter *TaskInputter) Name() string {
	return "TaskInputter"
}
