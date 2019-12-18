package middleware

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/service/redis"
	"gowatcher/go_monitor/utils"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetHeader(c, "token", "")
		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token is empty"))
			c.Abort()
			return
		}

		jwt := utils.NewJWT()
		claims, _ := jwt.ParseToken(token)

		//后面接口需要拿用户信息,todo:完善id机制
		c.Set("user_id", claims.UserName)

		isPass, _ := redis.QueryToken(c, token)
		if !isPass {
			c.JSON(http.StatusUnauthorized, utils.PackGinResult(http.StatusUnauthorized, "token error"))
			c.Abort()
			return
		}
	}
}
