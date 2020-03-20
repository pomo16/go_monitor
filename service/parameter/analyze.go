package parameter

import (
	"context"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//ParseAnalyzePolarityParams 解析情感分析参数
func ParseAnalyzePolarityParams(ctx context.Context, inputParams *model.InputParameter) (*model.AnalyzeParams, error) {
	if inputParams == nil {
		return nil, exceptions.ErrRequestParams
	}

	params := &model.AnalyzeParams{
		Content: inputParams.Content,
	}

	if params.Content == "" || len(params.Content) > 1000 {
		return nil, exceptions.ErrRequestParams
	}

	return params, nil
}
