package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/service/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	redis.InitRedis()
	redis.PingRedis()
}

func TestQueryToken(t *testing.T) {
	redis.InitRedis()
	ctx := gin.Context{}
	ctx.Set("user_name", "pomo")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InBvbW8iLCJleHAiOjE1NzY2NjA5MTMsImlzcyI6InBvbW8iLCJuYmYiOjE1NzY2NTkwMTN9.xtgDHxL1CVekpCUKxyn3S3tlQ21DctPDC3JlCBU5wP8"
	b, err := redis.QueryToken(&ctx, token)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("res:", b)
}

func TestIsTokenExisted(t *testing.T) {
	redis.InitRedis()
	ctx := gin.Context{}
	ctx.Set("user_name", "pomo1")
	res, err := redis.IsTokenExisted(&ctx)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(res)
}
