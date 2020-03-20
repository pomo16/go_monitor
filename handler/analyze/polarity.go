package analyze

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/processor"
	"gowatcher/go_monitor/processor/analyze"
)

//Polarity 获取情感分析结果
func Polarity(c *gin.Context) {
	context := model.NewAnalyzePolarityContext()

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

	analyzePolarityContextLoader := &analyze.PolarityLoader{}
	analyzePolarityContextCode := processor.LoaderCommon(c, context, analyzePolarityContextLoader)
	if analyzePolarityContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(analyzePolarityContextCode)
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
		"data":     packAnalyzePolarity(c, context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packAnalyzePolarity(c *gin.Context, context model.IAnalyzePolarityContext) map[string]interface{} {
	analyzePolarity := context.GetAnalyzePolarity()
	result := map[string]interface{}{
		"content":  analyzePolarity.Content,
		"polarity": analyzePolarity.Polarity,
		"score":    analyzePolarity.Score,
	}

	return result
}
