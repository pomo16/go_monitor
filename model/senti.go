package model

//SentimentInfo 情感信息
type SentimentInfo struct {
	Polarity string  `json:"polarity"` //极性标签
	Score    float64 `json:"score"`    //极性分数
}
