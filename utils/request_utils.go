package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

//GetParamString 解析字符串参数
func GetParamString(ctx *gin.Context, param string, defaultValue string) string {
	val := ctx.Request.FormValue(param)
	if val == "" {
		return defaultValue
	}
	return val
}

//GetParamInt32 解析int32参数
func GetParamInt32(ctx *gin.Context, param string, defalt int32) int32 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int32(rvl)
}

//GetParamInt16 解析int16参数
func GetParamInt16(ctx *gin.Context, param string, defalt int16) int16 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int16(rvl)
}

//GetParamInt64 解析int64参数
func GetParamInt64(ctx *gin.Context, param string, defalt int64) int64 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int64(rvl)
}

//GetParamInt 解析int参数
func GetParamInt(ctx *gin.Context, param string, defalt int) int {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return rvl
}

//PackGinResult 打包返回结果
func PackGinResult(code int, msg string) gin.H {
	return gin.H{
		"status_code": code,
		"msg":         msg,
	}
}

//AIDsSplit 对输入APPID字符串进行分割
func AIDsSplit(src string) []int64 {
	result := []int64{}
	str := strings.Split(src, "|")
	for _, s := range str {
		i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err != nil || i == 0 {
			continue
		}
		result = append(result, i)
	}
	return result
}

//GetHeader 获取头部
func GetHeader(ctx *gin.Context, param string, defaultValue string) string {
	value := ctx.Request.Header.Get(param)
	if value == "" {
		return defaultValue
	}
	return value
}

//SetHeader 设置头部
func SetHeader(ctx *gin.Context, param string, val string) {
	header := ctx.Writer.Header()
	header[param] = []string{val}
}
