package consts

import "time"

//log文件相关
const (
	LogFilePath = "go_monitor_log"
	LogFileName = "monitor"
)

//config文件路径
const ConfigFile = "config/config.yaml"

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

//修改方式
const (
	AddType    = 1 //添加模式
	UpdateType = 2 //更新模式
)

//消息响应
const (
	MsgSuccess = "success"
	MsgError   = "error"
)

//字符串时间类型
const (
	SQLTFormat = "2006-01-02 15:04:05" //SQL标准时间字符串格式
)

//Token
const (
	TokenHeader  = "X-Jwt-Token"
	TokenExpired = 15 * time.Minute
	Secret       = "sa9eu89pum94"
)

//Redis
const (
	RedisTokenPrefix = "monitor:tk:"
)

//Context
const (
	CtxUNameField = "user_name"
	CtxUIDField   = "user_id"
)

//Salt
const (
	UserIDSalt   = "uid%monitor_1"
	PasswordSalt = "pd%monitor_2"
)

//ES
const (
	ESTempIndex = "temp"
	MaxCount    = 20000
)
