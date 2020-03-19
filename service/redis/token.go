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

//IsTokenExisted 查看token是否存在，如果存在则返回现有token，并刷新token时间
func IsTokenExisted(c *gin.Context) (string, bool, error) {
	userName, ok := c.Get(consts.CtxUNameField)
	if !ok {
		return "", false, exceptions.ErrRedisHandle
	}
	key := consts.RedisTokenPrefix + userName.(string)
	token, err := redisClient.Get(key).Result()
	if err != nil {
		return "", false, exceptions.ErrRedisHandle
	}
	err = redisClient.Set(key, token, consts.TokenExpired).Err()
	if err != nil {
		return "", false, exceptions.ErrRedisHandle
	}
	return token, true, nil
}

//QueryToken 查询token
func QueryToken(c *gin.Context, token string) (bool, error) {
	userName, ok := c.Get(consts.CtxUNameField)
	if !ok {
		return false, exceptions.ErrRedisHandle
	}
	key := consts.RedisTokenPrefix + userName.(string)
	check, err := redisClient.Get(key).Result()
	if err != nil {
		return false, exceptions.ErrRedisHandle
	}

	return check == token, nil
}

//RemoveToken 删除token
func RemoveToken(c *gin.Context) {
	userName, _ := c.Get(consts.CtxUNameField)
	key := consts.RedisTokenPrefix + userName.(string)
	redisClient.Del(key)
}
