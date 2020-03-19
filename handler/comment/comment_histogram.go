package comment

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/processor"
	"gowatcher/go_monitor/processor/comment"
	"time"
)

//Histogram 获取评论直方图数据
func Histogram(c *gin.Context) {
	context := model.NewCommentHistoContext()

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

	commentHistoContextLoader := &comment.HistoLoader{}
	commentHistoContextCode := processor.LoaderCommon(c, context, commentHistoContextLoader)
	if commentHistoContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(commentHistoContextCode)
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
		"data":     packCommentHistogram(c, context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packCommentHistogram(c *gin.Context, context model.ICommentHistoContext) map[string]interface{} {
	commentHisto := context.GetCommentHisto()
	params := context.GetInputParameter()

	maxVal := int64(-1)
	times := make([]string, len(commentHisto.Counts))
	timeStep := consts.OneDay

	for idx, val := range commentHisto.Counts {
		times[idx] = time.Unix(params.BeginTime+int64(idx)*timeStep, 0).Format(consts.HISTOTFormat)
		if val > maxVal {
			maxVal = val
		}
	}

	result := map[string]interface{}{
		"begin_time": commentHisto.BeginTime,
		"end_time":   commentHisto.EndTime,
		"counts":     commentHisto.Counts,
		"times":      times,
	}
	return result
}
