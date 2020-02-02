package parameter

import (
	"context"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//checkOffset 检查搜索偏移量
func checkOffset(offset int) int {
	if offset < 0 {
		return 0
	}
	return offset
}

//checkLimit 检查数据查询条数
func checkLimit(limit int) int {
	if limit > consts.MaxCount {
		return consts.MaxCount
	} else if limit < 0 {
		return 10
	}

	return limit
}

//checkTimeSwitch 检查时间范围开关
func checkTimeSwitch(beginTime int64, endTime int64) bool {
	return beginTime != 0 && endTime != 0 && beginTime <= endTime
}

//checkAppIDSwitch 检查产品开关
func checkAppIDSwitch(aIDs []int64) bool {
	return len(aIDs) > 0
}

//ParseCommentListParams 解析查询参数
func ParseCommentListParams(ctx context.Context, inputParams *model.InputParameter) (*model.CommentListParams, error) {
	if inputParams == nil {
		return nil, exceptions.ErrValueEmpty
	}

	params := &model.CommentListParams{
		BeginTime: inputParams.BeginTime,
		EndTime:   inputParams.EndTime,
		OffSet:    inputParams.OffSet,
		Limit:     inputParams.Limit,
		MainID:    inputParams.MainID,
		QueryType: inputParams.CommentParams.QueryType,
		AIDs:      inputParams.AIDs,
	}

	if params.QueryType == consts.IdType && params.MainID == "" {
		return nil, exceptions.ErrValueEmpty
	}

	//时间错误，可以均为空
	if params.BeginTime > params.EndTime {
		return nil, exceptions.ErrTimeParams
	}

	params.OffSet = checkOffset(params.OffSet)
	params.Limit = checkLimit(params.Limit)
	params.TimeEnable = checkTimeSwitch(params.BeginTime, params.EndTime)
	params.APPIDEnable = checkAppIDSwitch(params.AIDs)
	return params, nil
}
