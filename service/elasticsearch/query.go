package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"strconv"
	"time"
)

//TimeStrToTimeStamp 时间字符串转时间戳
func TimeStrToTimeStamp(tStr string) int64 {
	tm, _ := time.Parse(consts.SQLTFormat, tStr)
	return tm.Unix()
}

//CommentCount 根据指定条件获取计数
func CommentCount(ctx context.Context, params *model.CommentCountParams) (*model.CommentCount, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = PublishTimeFilter(boolQuery, params.BeginTime, params.EndTime)

	if params.APPIDEnable {
		boolQuery = AIDsMatcher(boolQuery, params.AIDs)
	}

	if params.PolarityEnable {
		boolQuery = PolarityFilter(boolQuery, params.Polarity)
	}

	if params.KeywordEnable {
		boolQuery = KeywordMatcher(boolQuery, params.Keyword)
	}

	result, err := elasticClient.Count().
		Index(consts.ESIndex).
		Query(boolQuery).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get count error: %v", err)
		return nil, err
	}

	cnt := &model.CommentCount{
		BeginTime: params.BeginTime,
		EndTime:   params.EndTime,
		Count:     result,
	}

	return cnt, nil
}

//CommentHistogram 根据指定条件获取直方图数据
func CommentHistogram(ctx context.Context, params *model.CommentHistoParams) (*model.CommentHistogram, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = PublishTimeFilter(boolQuery, params.BeginTime, params.EndTime)

	if params.PolarityEnable {
		boolQuery = PolarityFilter(boolQuery, params.Polarity)
	}

	dateHistoAgg := elastic.NewDateHistogramAggregation().Field("publish_time").Interval("day").TimeZone("+08:00").MinDocCount(0)

	result, err := elasticClient.Search().
		Index(consts.ESIndex).
		Query(boolQuery).
		Size(0).
		Aggregation("HistoCount", dateHistoAgg).
		Do(ctx)

	if err != nil {
		logrus.Warnf("agg time error:%v", err)
		return nil, err
	}

	var bTime, eTime int64
	counts := []int64{}

	histogram, found := result.Aggregations.DateHistogram("HistoCount")
	if found {
		for _, val := range histogram.Buckets {
			counts = append(counts, val.DocCount)
		}

		nums := len(histogram.Buckets)
		if nums > 0 {
			bTime = TimeStrToTimeStamp(*(histogram.Buckets[0]).KeyAsString)
			eTime = TimeStrToTimeStamp(*(histogram.Buckets[nums-1]).KeyAsString)
		}
	}

	dateHisto := &model.CommentHistogram{
		BeginTime: bTime,
		EndTime:   eTime,
		Counts:    counts,
	}

	return dateHisto, nil
}

//CommentList 获取评论
func CommentList(ctx context.Context, params *model.CommentListParams) ([]*model.Comment, error) {
	boolQuery := elastic.NewBoolQuery()

	if params.TimeEnable {
		boolQuery = PublishTimeFilter(boolQuery, params.BeginTime, params.EndTime)
	}

	if params.APPIDEnable {
		boolQuery = AIDsMatcher(boolQuery, params.AIDs)
	}

	if params.KeywordEnable {
		boolQuery = KeywordMatcher(boolQuery, params.Keyword)
	}

	result, err := elasticClient.Search().
		Index(consts.ESIndex).
		Query(boolQuery).
		Sort("publish_timestamp", false).
		From(params.OffSet).
		Size(params.Limit).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get comments error: %v", err)
		return nil, err
	}

	commentList, err := commentOutputter(result)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}

//commentOutputter 评论输出
func commentOutputter(res *elastic.SearchResult) ([]*model.Comment, error) {
	commentList := []*model.Comment{}
	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			comment := &model.Comment{}
			if err := json.Unmarshal(hit.Source, comment); err != nil {
				return nil, exceptions.ErrParseResult
			}
			comment.Score, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", comment.Score), 64)
			commentList = append(commentList, comment)
		}
	}
	return commentList, nil
}

//QueryByMainID 通过MainID获取
func QueryByMainID(ctx context.Context, mainID string) ([]*model.Comment, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = MainIDFilter(boolQuery, mainID)

	result, err := elasticClient.Search().
		Index(consts.ESIndex).
		Query(boolQuery).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get comments error: %v", err)
		return nil, err
	}

	commentList, err := commentOutputter(result)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}
