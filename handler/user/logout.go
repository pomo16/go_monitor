package user

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/service/redis"
	"net/http"
)

//Logout 登出
func Logout(c *gin.Context) {
	redis.RemoveToken(c)
	c.JSON(http.StatusOK, gin.H{
		"message":  consts.MsgSuccess,
		"data":     map[string]interface{}{},
		"err_no":   0,
		"err_tips": "成功",
	})
}
