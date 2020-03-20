package model

type AnalyzeParams struct {
	Content string
}

type AnalyzePolarityParams struct {
	Content string
}

type AnalyzePolarity struct {
	Content  string  `json:"content"`
	Polarity string  `json:"polarity"`
	Score    float64 `json:"score"`
}
