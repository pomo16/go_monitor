package crawl

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
)

//TaskLoader 爬虫任务加载器
type TaskLoader struct{}

//Process 加载爬虫任务
func (loader *TaskLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	inputParameter := runCtx.GetInputParameter()

	switch inputParameter.QueryType {
	case consts.IdType:
		if inputParameter.TaskID == 0 {
			return exceptions.ErrRequestParams
		} else {
			if err := database.GetTask(ctx, inputParameter.TaskID); err == nil {
				return nil
			}
		}
	case consts.ListType:
		if err := database.ListTask(ctx); err == nil {
			return nil
		}
	default:
		return exceptions.ErrRequestParams
	}

	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *TaskLoader) Name() string {
	return "TaskLoader"
}
