package user

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service"
	"gowatcher/go_monitor/service/redis"
	"gowatcher/go_monitor/utils"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	parameter := service.ParseInputParameter(c)

	if parameter.UserName == "" || parameter.Password == "" {
		errCode := exceptions.ErrRequestParams
		logrus.Error("user %v login error: %v\n", parameter.UserName, errCode)
		errNo, errTips := exceptions.ErrConvert(errCode)
		c.JSON(http.StatusOK, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
	}

	//TODO 鉴权逻辑
	if parameter.UserName == "pomo" && parameter.Password == "123" {
		params := &model.LoginParams{
			UserName: parameter.UserName,
			Password: parameter.Password,
		}
		c.Set("user_id", params.UserName)
		generateToken(c, params)
		c.JSON(http.StatusOK, gin.H{
			"message":  consts.MsgSuccess,
			"data":     map[string]interface{}{},
			"err_no":   0,
			"err_tips": "成功",
		})
	} else {
		errCode := exceptions.ErrLogin
		logrus.Error("user %v login error: %v\n", parameter.UserName, errCode)
		errNo, errTips := exceptions.ErrConvert(errCode)
		c.JSON(http.StatusOK, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, params *model.LoginParams) {
	j := &utils.JWT{
		SigningKey: []byte(consts.Secret),
	}
	claims := utils.CustomClaims{
		params.UserName,
		jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 900,  // 过期时间15分钟
			Issuer:    "pomo",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	err = redis.SetToken(c, token)
	if err != nil {
		panic(err)
	}
	logrus.Info(token)

	return
}
