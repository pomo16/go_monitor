package model

//CrawlParams 爬虫任务参数集
type CrawlParams struct {
	TaskID    int32
	AppID     string
	AppName   string
	Status    int16
	QueryType int16
}

//CrawlTask 爬虫任务结构体
type CrawlTask struct {
	TaskID     int32
	AppID      string
	AppName    string
	Status     int16
	CreateTime string
	ModifyTime string
}
