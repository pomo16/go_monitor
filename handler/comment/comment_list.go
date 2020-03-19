package comment

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/processor"
	"gowatcher/go_monitor/processor/comment"
	"strconv"
)

//List 获取评论
func List(c *gin.Context) {
	context := model.NewCommentListContext()

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

	commentListContextLoader := &comment.ListLoader{}
	commentListContextCode := processor.LoaderCommon(c, context, commentListContextLoader)
	if commentListContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(commentListContextCode)
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
		"data":     packCommentList(context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packCommentList(context model.ICommentListContext) map[string]interface{} {
	commentList := context.GetCommentList()
	inputParameter := context.GetInputParameter()
	listMap := make([]map[string]interface{}, len(commentList))
	if inputParameter.CommentParams.QueryType == consts.IdType {
		rating, _ := strconv.ParseInt(commentList[0].Rating, 10, 64)
		listMap[0] = map[string]interface{}{
			"comment_id":   commentList[0].CommentId,
			"main_id":      commentList[0].MainId,
			"app_name":     commentList[0].AppName,
			"app_id":       commentList[0].AppID,
			"title":        commentList[0].Title,
			"content":      commentList[0].Content,
			"rating":       rating,
			"version":      commentList[0].Version,
			"polarity":     commentList[0].Polarity,
			"score":        commentList[0].Score,
			"publish_time": commentList[0].PublishTime,
			"crawl_time":   commentList[0].CrawlTime,
		}
	} else if inputParameter.CommentParams.QueryType == consts.ListType {
		for key, val := range commentList {
			rating, _ := strconv.ParseInt(commentList[key].Rating, 10, 64)
			listMap[key] = map[string]interface{}{
				"comment_id":   val.CommentId,
				"main_id":      val.MainId,
				"app_name":     val.AppName,
				"app_id":       val.AppID,
				"title":        val.Title,
				"content":      val.Content,
				"rating":       rating,
				"version":      val.Version,
				"polarity":     val.Polarity,
				"score":        val.Score,
				"publish_time": val.PublishTime,
				"crawl_time":   val.CrawlTime,
			}
		}
	}

	result := map[string]interface{}{
		"comment_list": listMap,
	}

	return result
}
