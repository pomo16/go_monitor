package test

import (
	"gowatcher/go_monitor/service/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	redis.InitRedis()
	redis.PingRedis()
}
