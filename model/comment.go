package model

type CommentParams struct {
	BeginTime int64 //开始时间
	EndTime   int64 //结束时间
	QueryType int16 //查询方式
	Polarity  string
	AIDs      []int64
	MainID    string
	OffSet    int
	Limit     int
	Keyword   string
}

type CommentCountParams struct {
	BeginTime      int64
	EndTime        int64
	Polarity       string
	PolarityEnable bool
	AIDs           []int64
	APPIDEnable    bool
	Keyword        string
	KeywordEnable  bool
}

type CommentHistoParams struct {
	BeginTime      int64
	EndTime        int64
	Polarity       string
	PolarityEnable bool
}

type CommentListParams struct {
	BeginTime     int64
	EndTime       int64
	TimeEnable    bool
	AIDs          []int64
	APPIDEnable   bool
	QueryType     int16
	MainID        string
	OffSet        int
	Limit         int
	Keyword       string
	KeywordEnable bool
}

type Comment struct {
	CommentId        string  `json:"comment_id"`
	MainId           string  `json:"main_id"`
	AppID            string  `json:"app_id"`
	AppName          string  `json:"app_name"`
	Title            string  `json:"title"`
	Content          string  `json:"content"`
	Rating           string  `json:"rating"`
	Version          string  `json:"version"`
	Polarity         string  `json:"polarity"`
	Score            float64 `json:"score"`
	PublishTime      string  `json:"publish_time"`
	PublishTimeStamp int64   `json:"publish_timestamp"`
	CrawlTime        string  `json:"crawl_time"`
	CrawlTimeStamp   int64   `json:"crawl_timestamp"`
}

type CommentCount struct {
	BeginTime int64 //开始时间
	EndTime   int64 //结束时间
	Count     int64 //计数
}

type CommentHistogram struct {
	BeginTime int64   //起始时间
	EndTime   int64   //结束时间
	Counts    []int64 //指定范围内数据量
}
