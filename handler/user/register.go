package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/service/parameter"
	"gowatcher/go_monitor/utils"
	"net/http"
)

//Register 注册
func Register(c *gin.Context) {
	parameter := parameter.ParseInpuctParameter(c)
	if parameter.UserName == "" || parameter.Password == "" {
		errCode := exceptions.ErrRequestParams
		logrus.Errorf("user %v register error: %v", parameter.UserName, errCode)
		errNo, errTips := exceptions.ErrConvert(errCode)
		c.JSON(http.StatusOK, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	params := &model.LoginParams{
		UserName: parameter.UserName,
		Password: utils.Md5AddSalt(parameter.Password, consts.PasswordSalt, false),
	}

	err := database.InsertUser(c, params)
	if err != nil {
		logrus.Errorf("user %v register error: %v", params.UserName, err)
		c.JSON(http.StatusOK, utils.PackGinResult(http.StatusUnauthorized, "register error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     consts.MsgSuccess,
		"status_code": 200,
		"user":        params.UserName,
	})

}
