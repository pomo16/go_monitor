package redis

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
)

//SetToken 保存token
func SetToken(c *gin.Context, token string) error {
	userName, ok := c.Get(consts.CtxUNameField)
	if !ok {
		return exceptions.ErrRedisHandle
	}
	key := consts.RedisTokenPrefix + userName.(string)
	err := redisClient.Set(key, token, consts.TokenExpired).Err()
	if err != nil {
		return exceptions.ErrRedisHandle
	}
	return nil
}

//QueryToken 查询token
func QueryToken(c *gin.Context, token string) (bool, error) {
	userID, ok := c.Get(consts.CtxUNameField)
	if !ok {
		return false, exceptions.ErrRedisHandle
	}
	key := consts.RedisTokenPrefix + userID.(string)
	check, err := redisClient.Get(key).Result()
	if err != nil {
		return false, exceptions.ErrRedisHandle
	}

	return check == token, nil
}

//RemoveToken 删除token
func RemoveToken(c *gin.Context) {
	userID, _ := c.Get(consts.CtxUNameField)
	key := consts.RedisTokenPrefix + userID.(string)
	redisClient.Del(key)
}
