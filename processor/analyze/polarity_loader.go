package analyze

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/parameter"
	"gowatcher/go_monitor/service/text_analyze"
	"strconv"
)

//PolarityLoader 情感分析loader
type PolarityLoader struct{}

//Process 获取情感分析
func (loader *PolarityLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	cntCtx, ok := runCtx.(model.IAnalyzePolarityContext)
	if !ok {
		logrus.Warn(ctx, "Analyze Polarity loader context error")
		return exceptions.ErrTypeAssert
	}

	inputParameter := runCtx.GetInputParameter()

	polParams, err := parameter.ParseAnalyzePolarityParams(ctx, inputParameter)
	if err != nil {
		logrus.Error("Analyze Polarity loader params error")
		return err
	}

	senti, err := text_analyze.GetSentimentPolarity(polParams.Content)

	if senti != nil {
		senti.Score, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", senti.Score), 64)
		analyzePolarity := &model.AnalyzePolarity{
			Content:  polParams.Content,
			Polarity: senti.Polarity,
			Score:    senti.Score,
		}
		cntCtx.SetAnalyzePolarity(analyzePolarity)
		return nil
	}

	if err == nil && senti == nil {
		return exceptions.ErrResultEmpty
	}

	logrus.Error("Analyze Polarity loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *PolarityLoader) Name() string {
	return "PolarityLoader"
}
