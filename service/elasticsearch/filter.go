package elasticsearch

import "github.com/olivere/elastic/v7"

//PublishTimeFilter 返回评论时间范围条件时间戳比对
func PublishTimeFilter(boolQ *elastic.BoolQuery, bTime int64, eTime int64) *elastic.BoolQuery {
	return boolQ.Filter(elastic.NewRangeQuery("publish_timestamp").Gte(bTime).Lte(eTime))
}

//MainIDFilter ID过滤器
func MainIDFilter(boolQ *elastic.BoolQuery, mainID string) *elastic.BoolQuery {
	return boolQ.Must(elastic.NewTermQuery("main_id", mainID))
}

//AIDsMatcher 多APPID匹配
func AIDsMatcher(boolQ *elastic.BoolQuery, aIDs []int64) *elastic.BoolQuery {
	boolSubQ := elastic.NewBoolQuery()
	for _, w := range aIDs {
		boolSubQ.Should(elastic.NewTermQuery("aid", w))
	}
	return boolQ.Must(boolSubQ.MinimumNumberShouldMatch(1)) //至少匹配一个APPID
}

//PolarityFilter 情感过滤条件
func PolarityFilter(boolQ *elastic.BoolQuery, polarity string) *elastic.BoolQuery {
	return boolQ.Must(elastic.NewMatchQuery("polarity", polarity))
}
