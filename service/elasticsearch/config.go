package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"os"
	"path/filepath"
)

var (
	elasticClient *elastic.Client
)

//InitElasticSearch 初始化ES实例
func InitElasticSearch() {
	esAddr, err := ReadYamlConfig()
	if err != nil {
		panic(err)
	}

	elasticClient, err = elastic.NewClient(
		elastic.SetURL(esAddr),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	//需要看下到es服务端是否正常
	info, code, err := elasticClient.Ping(esAddr).Do(context.Background())
	if err != nil {
		panic(err)
	}

	logrus.Info("ElasticSearch returned with code ", code, " and version ", info.Version.Number)
}

//ReadYamlConfig 读取yaml配置文件返回ES链接
func ReadYamlConfig() (string, error) {
	path, _ := filepath.Abs(consts.ConfigFile)
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return "", exceptions.ErrConfigRead
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}

	esConfig := conf.ElasticSearch
	link := "http://" + esConfig.Host + ":" + esConfig.Port
	return link, nil
}
