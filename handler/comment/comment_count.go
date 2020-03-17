package comment

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/processor"
	"gowatcher/go_monitor/processor/comment"
)

//Count 获取评论计数
func Count(c *gin.Context) {
	context := model.NewCommentCountContext()

	preContextLoader := &processor.PreContextLoader{}
	preContextCode := processor.LoaderCommon(c, context, preContextLoader)
	if preContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(preContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	commentCountContextLoader := &comment.CountLoader{}
	commentCountContextCode := processor.LoaderCommon(c, context, commentCountContextLoader)
	if commentCountContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(commentCountContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  consts.MsgSuccess,
		"data":     packCommentCount(c, context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packCommentCount(c *gin.Context, context model.ICommentCountContext) map[string]interface{} {
	commentCount := context.GetCommentCount()
	result := map[string]interface{}{
		"count":      commentCount.Count,
		"begin_time": commentCount.BeginTime,
		"end_time":   commentCount.EndTime,
	}

	return result
}
