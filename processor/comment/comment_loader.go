package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/elasticsearch"
	"gowatcher/go_monitor/service/parameter"
)

//CommentLoader 评论加载器
type CommentLoader struct{}

//Process 加载评论
func (loader *CommentLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	listCtx, ok := runCtx.(model.ICommentListContext)
	if !ok {
		logrus.Warn(ctx, "CommentList loader listCtx error")
		return exceptions.ErrTypeAssert
	}

	var commentList []*model.Comment
	var err error
	inputParameter := runCtx.GetInputParameter()
	params, err := parameter.ParseCommentListParams(ctx, inputParameter)
	if err != nil {
		return err
	}

	switch params.QueryType {
	case consts.IdType:
		commentList, err = elasticsearch.QueryByMainID(ctx, params.MainID)
	case consts.ListType:
		commentList, err = elasticsearch.CommentList(ctx, params)
	default:
		return exceptions.ErrRequestParams
	}

	if commentList != nil && len(commentList) != 0 {
		listCtx.SetCommentList(commentList)
		return nil
	}

	if len(commentList) == 0 {
		return exceptions.ErrResultEmpty
	}

	logrus.Error("CommentList loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *CommentLoader) Name() string {
	return "CommentLoader"
}
