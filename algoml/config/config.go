package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//buildEnvPath 构造模型文件的根目录
func buildEnvPath(sub string) string {
	path, err := filepath.Abs(sub)
	if err != nil {
		panic(err)
	}
	return path
}

//模型运行时的资源目录
var (
	DataDir  = buildEnvPath("data")
	DictDir  = buildEnvPath("data/dict")
	ConfDir  = buildEnvPath("conf")
	TrainDir = buildEnvPath("data/train")
	ModelDir = buildEnvPath("data/model")
)

//TopicConfig 模型配置类型
type TopicConfig struct {
	Type          string  `json:"type"`
	NumTopics     int32   `json:"num_topics"`
	Alpha         float32 `json:"alpha"`
	Beta          float32 `json:"beta"`
	WordTopicFile string  `json:"word_topic_file"`
	VocabFile     string  `json:"vocab_file"`
	TweModelFile  string  `json:"twe_model_file,omitempty"`
}

//LoadTopicConfig 加载主题模型配置文件
func LoadTopicConfig(subDir string, fileName string) *TopicConfig {
	filePath := GetModelPath(subDir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	var conf TopicConfig
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil
	}

	conf.WordTopicFile = GetModelPath(subDir, conf.WordTopicFile)
	conf.VocabFile = GetModelPath(subDir, conf.VocabFile)
	return &conf
}

//GetModelPath 根据模型指定目录与指定文件获取模型路径
func GetModelPath(modelSubDir string, fileName string) string {
	modelFile := modelSubDir + "/" + fileName
	return path.Join(ModelDir, modelFile)
}
