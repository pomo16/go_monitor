package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/elasticsearch"
)

//CountLoader 评论信息计数loader
type CountLoader struct{}

//Process 获取评论计数数据
func (loader *CountLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	cntCtx, ok := runCtx.(model.ICommentCountContext)
	if !ok {
		logrus.Warn(ctx, "Comment Count loader context error")
		return exceptions.ErrTypeAssert
	}

	inputParameter := runCtx.GetInputParameter()

	//时间错误
	if inputParameter.BeginTime == 0 || inputParameter.EndTime == 0 || inputParameter.BeginTime > inputParameter.EndTime {
		logrus.Error("Comment Count loader params error")
		return exceptions.ErrTimeParams
	}

	cntParams := &model.CommentCountParams{
		BeginTime: inputParameter.BeginTime,
		EndTime:   inputParameter.EndTime,
		AIDs:      inputParameter.AIDs,
	}

	cntInfo, err := elasticsearch.CommentCount(ctx, cntParams)
	if cntInfo != nil {
		cntCtx.SetCommentCount(cntInfo)
		return nil
	}

	if err == nil && cntInfo == nil {
		return exceptions.ErrResultEmpty
	}

	logrus.Error("Comment Count loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *CountLoader) Name() string {
	return "CountLoader"
}
