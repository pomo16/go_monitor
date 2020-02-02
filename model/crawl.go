package model

//CrawlParams 爬虫任务参数集
type CrawlParams struct {
	TaskID     int32
	AppID      string
	AppName    string
	Status     int16
	QueryType  int16 //查询方式
	ConfigType int16 //配置方式：新增/更新
}

//CrawlTask 爬虫任务结构体
type CrawlTask struct {
	ID         int32  `json:"id" gorm:"column:id"`
	AppID      string `json:"app_id" gorm:"column:app_id"`
	AppName    string `json:"app_name" gorm:"column:app_name"`
	Status     int16  `json:"status" gorm:"column:status"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
	ModifyTime string `json:"modify_time" gorm:"column:modify_time"`
}
