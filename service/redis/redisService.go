package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
)

func PingRedis() {
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

func SetToken(c *gin.Context, token string) error {
	userID, ok := c.Get("user_id")
	if !ok {
		return exceptions.ErrRedisHandle
	}
	key := consts.RedisTokenPrefix + userID.(string)
	err := redisClient.Set(key, token, consts.TokenExpired).Err()
	if err != nil {
		return exceptions.ErrRedisHandle
	}
	return nil
}

func QueryToken(c *gin.Context, token string) (bool, error) {
	userID, ok := c.Get("user_id")
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
