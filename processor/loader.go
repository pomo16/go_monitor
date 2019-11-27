package processor

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//IProcessor 处理器接口
type IProcessor interface {
	Process(*gin.Context, model.IContext) exceptions.ErrProcessor
	Name() string
}

//LoaderCommon 加载通用任务处理
func LoaderCommon(ctx *gin.Context, runCtx model.IContext, processor IProcessor) exceptions.ErrProcessor {
	code := exceptions.ErrProcessPanic
	code = processor.Process(ctx, runCtx)
	return code
}
