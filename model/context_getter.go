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

//SetAddTask 设置爬虫任务添加
func (infoCtx *AddTaskContext) SetAddTask(crawlTask *CrawlTask) {
	infoCtx.crawlTask = crawlTask
}

//GetAddTask 获取爬虫任务添加
func (infoCtx *AddTaskContext) GetAddTask() *CrawlTask {
	if infoCtx.crawlTask != nil {
		return infoCtx.crawlTask
	}
	return &CrawlTask{}
}
