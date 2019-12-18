package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/service/redis"
	"gowatcher/go_monitor/utils"
	"net/http"
)

func Refresh(c *gin.Context) {
	token := utils.GetHeader(c, "token", "")
	newToken, err := utils.NewJWT().RefreshToken(token)
	err = redis.SetToken(c, newToken)
	if err != nil {
		errCode := exceptions.ErrToken
		logrus.Error("refresh token error: %v\n", errCode)
		errNo, errTips := exceptions.ErrConvert(errCode)
		c.JSON(http.StatusOK, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  consts.MsgSuccess,
		"data":     newToken,
		"err_no":   0,
		"err_tips": "成功",
	})
}
