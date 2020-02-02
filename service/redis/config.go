package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"gopkg.in/yaml.v2"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"os"
	"path/filepath"
)

var (
	redisClient *redis.Client
)

func InitRedis() {
	redisOps, err := ReadYamlConfig()
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(redisOps)
}

//ReadYamlConfig 读取yaml配置文件返回redis配置
func ReadYamlConfig() (*redis.Options, error) {
	path, _ := filepath.Abs(consts.ConfigFile)
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return nil, exceptions.ErrConfigRead
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}

	redisConfig := conf.Redis
	return &redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	}, nil
}

//PingRedis 测试连接Redis
func PingRedis() {
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}
