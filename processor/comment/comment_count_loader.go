package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/parameter"
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

	cntParams, err := parameter.ParseCommentCountParams(ctx, inputParameter)
	if err != nil {
		logrus.Error("Comment Count loader params error")
		return err
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
