package middleware

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/service/redis"
	"gowatcher/go_monitor/utils"
	"net/http"
)

//CheckLogin 登录态校验
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetHeader(c, consts.TokenHeader, "")
		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token is empty"))
			c.Abort()
			return
		}

		jwt := utils.NewJWT()
		claims, _ := jwt.ParseToken(token)
		c.Set(consts.CtxUIDField, claims.UserID)
		c.Set(consts.CtxUNameField, claims.UserName)

		isPass, _ := redis.QueryToken(c, token)
		if !isPass {
			c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token error"))
			c.Abort()
			return
		}
	}
}
