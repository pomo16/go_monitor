package service

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/utils"
)

func ParseInputParameter(ctx *gin.Context) *model.InputParameter {
	parameter := &model.InputParameter{
		LoginParams: model.LoginParams{
			UserName: utils.GetParamString(ctx, "username", ""),
			Password: utils.GetParamString(ctx, "password", ""),
		},
		CrawlParams: model.CrawlParams{
			TaskID:     utils.GetParamInt32(ctx, "task_id", 0),
			AppID:      utils.GetParamString(ctx, "app_id", ""),
			AppName:    utils.GetParamString(ctx, "app_name", ""),
			Status:     utils.GetParamInt16(ctx, "status", 0),
			QueryType:  utils.GetParamInt16(ctx, "q_type", 0),
			ConfigType: utils.GetParamInt16(ctx, "c_type", 0),
		},
	}

	return parameter
}
