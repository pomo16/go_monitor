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

//checkKeywordSwitch 检查关键词开关
func checkKeywordSwitch(keyword string) bool {
	return keyword != ""
}

//checkPolaritySwitch 检查情感开关
func checkPolaritySwitch(polarity string) bool {
	return polarity != ""
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
		Keyword:   inputParams.Keyword,
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
	params.KeywordEnable = checkKeywordSwitch(params.Keyword)
	return params, nil
}

//ParseCommentCountParams 解析计数参数
func ParseCommentCountParams(ctx context.Context, inputParams *model.InputParameter) (*model.CommentCountParams, error) {
	if inputParams == nil {
		return nil, exceptions.ErrValueEmpty
	}

	params := &model.CommentCountParams{
		BeginTime: inputParams.BeginTime,
		EndTime:   inputParams.EndTime,
		Polarity:  inputParams.Polarity,
		AIDs:      inputParams.AIDs,
		Keyword:   inputParams.Keyword,
	}

	//时间错误
	if inputParams.BeginTime == 0 || inputParams.EndTime == 0 || inputParams.BeginTime > inputParams.EndTime {
		return nil, exceptions.ErrTimeParams
	}

	params.PolarityEnable = checkPolaritySwitch(params.Polarity)
	params.APPIDEnable = checkAppIDSwitch(params.AIDs)
	params.KeywordEnable = checkKeywordSwitch(params.Keyword)
	return params, nil
}

//ParseCommentHistoParams 解析直方图参数
func ParseCommentHistoParams(ctx context.Context, inputParams *model.InputParameter) (*model.CommentHistoParams, error) {
	if inputParams == nil {
		return nil, exceptions.ErrValueEmpty
	}

	params := &model.CommentHistoParams{
		BeginTime: inputParams.BeginTime,
		EndTime:   inputParams.EndTime,
		Polarity:  inputParams.Polarity,
	}

	//时间错误
	if inputParams.BeginTime == 0 || inputParams.EndTime == 0 || inputParams.BeginTime > inputParams.EndTime {
		return nil, exceptions.ErrTimeParams
	}

	params.PolarityEnable = checkPolaritySwitch(params.Polarity)
	return params, nil
}
