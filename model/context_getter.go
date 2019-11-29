package model

//SetInputParameter 设置输入参数
func (infoCtx *BaseContext) SetInputParameter(parameter *InputParameter) {
	infoCtx.inputParameter = parameter
}

//GetInputParameter 设置输出参数
func (infoCtx *BaseContext) GetInputParameter() *InputParameter {
	if infoCtx.inputParameter != nil {
		return infoCtx.inputParameter
	}
	return &InputParameter{}
}

//SetTaskConf 设置爬虫任务配置参数
func (infoCtx *TaskConfContext) SetTaskConf(taskConf *CrawlTask) {
	infoCtx.taskConf = taskConf
}

//GetTaskConf 获取爬虫任务配置参数
func (infoCtx *TaskConfContext) GetTaskConf() *CrawlTask {
	if infoCtx.taskConf != nil {
		return infoCtx.taskConf
	}
	return &CrawlTask{}
}

//SetGetTask 设置爬虫任务获取参数
func (infoCtx *TaskListContext) SetTaskList(taskList []*CrawlTask) {
	infoCtx.taskList = taskList
}

//GetAddTask 获取爬虫任务添加
func (infoCtx *TaskListContext) GetTaskList() []*CrawlTask {
	if infoCtx.taskList != nil {
		return infoCtx.taskList
	}
	return []*CrawlTask{}
}
