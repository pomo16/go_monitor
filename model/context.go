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

//TaskConfContext 爬虫任务配置上下文
type TaskConfContext struct {
	BaseContext
	taskConf *CrawlTask
}

//ITaskConfContext 爬虫任务配置接口
type ITaskConfContext interface {
	IContext
	SetTaskConf(taskConf *CrawlTask)
	GetTaskConf() *CrawlTask
}

//NewTaskConfContext context构造函数
func NewTaskConfContext() *TaskConfContext {
	return &TaskConfContext{}
}

//TaskListContext 爬虫任务获取上下文
type TaskListContext struct {
	BaseContext
	taskList []*CrawlTask
}

//ITaskListContext 爬虫任务获取接口
type ITaskListContext interface {
	IContext
	SetTaskList(taskList []*CrawlTask)
	GetTaskList() []*CrawlTask
}

//NewTaskListContext context构造函数
func NewTaskListContext() *TaskListContext {
	return &TaskListContext{}
}

//CommentListContext 评论获取上下文
type CommentListContext struct {
	BaseContext
	commentList []*Comment
}

//ICommentListContext 评论获取接口
type ICommentListContext interface {
	IContext
	SetCommentList(commentList []*Comment)
	GetCommentList() []*Comment
}

//NewCommentListContext context构造函数
func NewCommentListContext() *CommentListContext {
	return &CommentListContext{}
}
