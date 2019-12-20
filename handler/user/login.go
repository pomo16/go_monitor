package user

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service"
	"gowatcher/go_monitor/service/database"
	"gowatcher/go_monitor/service/redis"
	"gowatcher/go_monitor/utils"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	parameter := service.ParseInputParameter(c)

	if parameter.UserName == "" || parameter.Password == "" {
		errCode := exceptions.ErrRequestParams
		logrus.Errorf("user %v login error: %v", parameter.UserName, errCode)
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

	user, err := database.CheckUser(c, params)
	if err != nil {
		logrus.Errorf("user %v login error: %v", params.UserName, err)
		c.JSON(http.StatusOK, utils.PackGinResult(http.StatusUnauthorized, "login error"))
		return
	}

	if user != nil {
		c.Set(consts.CtxUIDField, user.UserID)
		c.Set(consts.CtxUNameField, user.UserName)
		generateToken(c)
	} else {
		logrus.Errorf("user %v login error: %v", params.UserName, exceptions.ErrLogin)
		c.JSON(http.StatusOK, utils.PackGinResult(http.StatusUnauthorized, "login error"))
	}
}

// 生成令牌
func generateToken(c *gin.Context) {
	j := &utils.JWT{
		SigningKey: []byte(consts.Secret),
	}

	var ok bool
	var userID, userName interface{}
	userID, ok = c.Get(consts.CtxUIDField)
	userName, ok = c.Get(consts.CtxUNameField)
	if !ok {
		c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token error"))
		logrus.Errorf("create token err")
		return
	}

	claims := utils.CustomClaims{
		UserID:   userID.(string),
		UserName: userName.(string),
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                   // 签名生效时间
			ExpiresAt: time.Now().Add(consts.TokenExpired).Unix(), // 过期时间15分钟
			Issuer:    "pomo",                                     //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token error"))
		logrus.Errorf("create token err: %v", err)
		return
	}

	err = redis.SetToken(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token error"))
		logrus.Errorf("save token err: %v", err)
		return
	}

	logrus.Info(token)
	c.JSON(http.StatusOK, utils.PackGinResult(http.StatusOK, "login success!"))

	return
}
