package text_analyze

import (
	"gowatcher/go_monitor/algoml"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//GetSentimentPolarity 调用情感分析模型分析文本极性
func GetSentimentPolarity(doc string) (*model.SentimentInfo, error) {
	if doc == "" {
		return nil, exceptions.ErrValueEmpty
	}

	score := algoml.SentiML.Classify(doc)
	polarity := EchoPolarity(score)
	result := &model.SentimentInfo{
		Score:    score,
		Polarity: polarity,
	}

	return result, nil
}

//EchoPolarity 将情感分析分数转换为极性标签
func EchoPolarity(score float64) string {
	polarity := ""
	if score < 0.2 {
		polarity = "neg"
	} else if score > 0.8 {
		polarity = "pos"
	} else {
		polarity = "net"
	}
	return polarity
}
