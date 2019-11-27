package processor

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service"
)

//PreContextLoader 请求预处理
type PreContextLoader struct{}

//Process 预处理过程
func (dif *PreContextLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	parameter := service.ParseInputParameter(ctx)
	runCtx.SetInputParameter(parameter)
	return nil
}

//Name Processor名称
func (dif *PreContextLoader) Name() string {
	return "PreContextLoader"
}
