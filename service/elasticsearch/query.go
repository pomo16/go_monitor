package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
)

//CommentCount 根据指定条件获取计数
func CommentCount(ctx context.Context, params *model.CommentCountParams) (*model.CommentCount, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = PublishTimeFilter(boolQuery, params.BeginTime, params.EndTime)

	result, err := elasticClient.Count().
		Index(consts.ESTempIndex).
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

//CommentList 获取评论
func CommentList(ctx context.Context, params *model.CommentListParams) ([]*model.Comment, error) {
	boolQuery := elastic.NewBoolQuery()

	if params.TimeEnable {
		boolQuery = PublishTimeFilter(boolQuery, params.BeginTime, params.EndTime)
	}

	if params.APPIDEnable {
		boolQuery = AIDsMatcher(boolQuery, params.AIDs)
	}

	result, err := elasticClient.Search().
		Index(consts.ESTempIndex).
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
		Index(consts.ESTempIndex).
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
