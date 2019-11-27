package model

//BaseContext 基础上下文信息
type BaseContext struct {
	inputParameter *InputParameter
}

//IContext 上下文结构读写接口
type IContext interface {
	SetInputParameter(parameter *InputParameter)
	GetInputParameter() *InputParameter
}

//AddTaskContext 爬虫任务添加上下文
type AddTaskContext struct {
	BaseContext
	crawlTask *CrawlTask
}

//IAddTaskContext 爬虫任务添加接口
type IAddTaskContext interface {
	IContext
	SetAddTask(crawlTask *CrawlTask)
	GetAddTask() *CrawlTask
}

//NewAddTaskContext context构造函数
func NewAddTaskContext() *AddTaskContext {
	return &AddTaskContext{}
}
