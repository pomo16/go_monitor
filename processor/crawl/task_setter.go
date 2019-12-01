package crawl

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
	"time"
)

//TaskSetter 爬虫任务配置器
type TaskSetter struct{}

//Process 配置爬虫任务
func (setter *TaskSetter) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	inputParameter := runCtx.GetInputParameter()
	if inputParameter.TaskID == 0 {
		return exceptions.ErrRequestParams
	}
	if inputParameter.Status < consts.Normal || inputParameter.Status > consts.Unused {
		return exceptions.ErrRequestParams
	}

	params := &model.CrawlTask{
		ID:         inputParameter.TaskID,
		AppID:      inputParameter.AppID,
		AppName:    inputParameter.AppName,
		Status:     inputParameter.Status,
		CreateTime: time.Now().Format(consts.SQLTFormat),
		ModifyTime: time.Now().Format(consts.SQLTFormat),
	}

	switch inputParameter.ConfigType {
	case consts.AddType:
		if err := database.InsertTask(ctx, params); err == nil {
			return nil
		}
	case consts.UpdateType:
		if err := database.UpdateTask(ctx, params); err == nil {
			return nil
		}
	default:
		return exceptions.ErrRequestParams
	}

	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (setter *TaskSetter) Name() string {
	return "TaskSetter"
}
