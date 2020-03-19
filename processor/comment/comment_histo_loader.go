package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/parameter"
)

//HistoLoader 评论直方图数据loader
type HistoLoader struct{}

//Process 获取评论直方图数据
func (loader *HistoLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	histoCtx, ok := runCtx.(model.ICommentHistoContext)
	if !ok {
		logrus.Warn(ctx, "Comment Histogram loader context error")
		return exceptions.ErrTypeAssert
	}

	inputParameter := runCtx.GetInputParameter()

	histoParams, err := parameter.ParseCommentHistoParams(ctx, inputParameter)
	if err != nil {
		logrus.Error("Comment Histogram loader params error")
		return err
	}

	histoInfo, err := elasticsearch.CommentHistogram(ctx, histoParams)
	if histoInfo != nil {
		histoCtx.SetCommentHisto(histoInfo)
		return nil
	}

	if err == nil && histoInfo == nil {
		return exceptions.ErrResultEmpty
	}

	logrus.Error("Comment Histogram loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *HistoLoader) Name() string {
	return "HistoLoader"
}
