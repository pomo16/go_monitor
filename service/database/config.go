package database

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"os"
	"path/filepath"
)

var (
	DataBase *gorm.DB
)

func InitDB() {
	dbLink, err := ReadYamlConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", dbLink)
	if err == nil {
		db.SingularTable(true)
		DataBase = db
	} else {
		panic(err)
	}
}

//ReadYamlConfig 读取yaml配置文件返回数据库链接
func ReadYamlConfig() (string, error) {
	path, _ := filepath.Abs(consts.ConfigFile)
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return "", exceptions.ErrConfigRead
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}

	dbConfig := conf.Mysql
	link := dbConfig.UserName + ":" + dbConfig.Password +
		"@tcp(" + dbConfig.Host + ":" + dbConfig.Port +
		")/gowatcher?charset=utf8&parseTime=True&loc=Local"
	return link, nil
}
