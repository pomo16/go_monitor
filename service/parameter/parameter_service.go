package parameter

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/utils"
)

func ParseInputParameter(ctx *gin.Context) *model.InputParameter {
	parameter := &model.InputParameter{
		LoginParams: utils.GetLoginParams(ctx),
		CrawlParams: model.CrawlParams{
			TaskID:     utils.GetParamInt32(ctx, "task_id", 0),
			AppID:      utils.GetParamString(ctx, "app_id", ""),
			AppName:    utils.GetParamString(ctx, "app_name", ""),
			Status:     utils.GetParamInt16(ctx, "status", 0),
			QueryType:  utils.GetParamInt16(ctx, "q_type", 0),
			ConfigType: utils.GetParamInt16(ctx, "c_type", 0),
		},
		CommentParams: model.CommentParams{
			BeginTime: utils.GetParamInt64(ctx, "begin_time", 0),
			EndTime:   utils.GetParamInt64(ctx, "end_time", 0),
			MainID:    utils.GetParamString(ctx, "main_id", ""),
			Polarity:  utils.GetParamString(ctx, "polarity", ""),
			QueryType: utils.GetParamInt16(ctx, "q_type", 0),
			Limit:     utils.GetParamInt(ctx, "limit", 10),
			OffSet:    utils.GetParamInt(ctx, "offset", 0),
			AIDs:      utils.AIDsSplit(utils.GetParamString(ctx, "aids", "")),
		},
	}

	return parameter
}
