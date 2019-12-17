package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParamString(ctx *gin.Context, param string, defaultValue string) string {
	val := ctx.Request.FormValue(param)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetParamInt32(ctx *gin.Context, param string, defalt int32) int32 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int32(rvl)
}

func GetParamInt16(ctx *gin.Context, param string, defalt int16) int16 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int16(rvl)
}

//PackGinResult 打包返回结果
func PackGinResult(code int, msg string) gin.H {
	return gin.H{
		"status_code": code,
		"msg":         msg,
	}
}

func GetHeader(ctx *gin.Context, param string, defaultValue string) string {
	value := ctx.Request.Header.Get(param)
	if value == "" {
		return defaultValue
	}
	return value
}

func SetHeader(ctx *gin.Context, param string, val string) {
	header := ctx.Writer.Header()
	header[param] = []string{val}
}
