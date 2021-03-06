package exceptions

import "errors"

//ErrProcessor 定义loader错误类型
type ErrProcessor error

//用错误类型标识异常种类
var (
	ErrConfigRead  = errors.New("read config error")
	ErrDBHandle    = errors.New("db handle failed")
	ErrRedisHandle = errors.New("redis handle failed")

	ErrResultEmpty   = errors.New("result empty")
	ErrRequestParams = errors.New("illegal request parameters")
	ErrTimeParams    = errors.New("time params error")
	ErrValueEmpty    = errors.New("value empty")

	ErrProcessPanic  = errors.New("processor panic")
	ErrProcessFailed = errors.New("processor failed")
	ErrTypeAssert    = errors.New("type assert error")

	ErrLogin = errors.New("login error")
	ErrToken = errors.New("token error")

	ErrParseResult = errors.New("result parse error")
)

//返回给前端的业务错误码err_no
const (
	Normal       int64 = 0
	SystemBusy   int64 = 3001
	IllegalParam int64 = 3002
	ResultEmpty  int64 = 3003
	AuthError    int64 = 3004
)

//ErrTips 将对应业务错误码的错误信息返回给前端
func ErrTips(errNo int64) string {
	var tips string
	switch errNo {
	case Normal:
		tips = "成功"
	case SystemBusy:
		tips = "系统繁忙"
	case IllegalParam:
		tips = "参数非法"
	case ResultEmpty:
		tips = "结果为空"
	case AuthError:
		tips = "鉴权失败"
	default:
		tips = "未知错误"
	}
	return tips
}

//ErrConvert 将系统错误转换为对应的errNo与errTips
func ErrConvert(err error) (int64, string) {
	errNo := Normal
	switch err {
	case ErrRequestParams:
		errNo = IllegalParam
	case ErrResultEmpty:
		errNo = ResultEmpty
	case ErrLogin:
		errNo = AuthError
	case ErrToken:
		errNo = AuthError
	default:
		errNo = SystemBusy
	}
	return errNo, ErrTips(errNo)
}
