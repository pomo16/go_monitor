package service

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/utils"
)

func ParseInputParameter(ctx *gin.Context) *model.InputParameter {
	parameter := &model.InputParameter{
		CrawlParams: model.CrawlParams{
			TaskID:    utils.GetParamInt32(ctx, "task_id", 0),
			AppID:     utils.GetParamString(ctx, "app_id", ""),
			AppName:   utils.GetParamString(ctx, "app_name", ""),
			Status:    utils.GetParamInt16(ctx, "status", 0),
			QueryType: utils.GetParamInt16(ctx, "type", 0),
		},
	}

	return parameter
}
