package consts

//log文件相关
const (
	LogFilePath = "output/monitor_log"
	LogFileName = "monitor"
)

//config文件路径
const ConfigFile = "../config/config.yaml"

//爬虫任务状态
const (
	Normal = 1 //启用中
	Unused = 2 //未启用
)

//查询方式
const (
	IdType   = 1 //按主键单条查找
	ListType = 2 //列表形式查询所有
)

//消息响应
const (
	MsgSuccess = "success"
	MsgError   = "error"
)
