package test

import (
	"context"
	"fmt"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/service/elasticsearch"
	"testing"
)

func TestES(t *testing.T) {
	elasticsearch.InitElasticSearch()
}

func TestCommentList(t *testing.T) {
	elasticsearch.InitElasticSearch()
	ctx := context.Background()
	//res, err := elasticsearch.QueryByMainID(ctx, "27ed3b0a86895e1731d37ddcf5202318")
	params := &model.CommentListParams{
		BeginTime:     1584288000,
		EndTime:       1584892800,
		TimeEnable:    true,
		AIDs:          nil,
		APPIDEnable:   false,
		QueryType:     2,
		MainID:        "",
		OffSet:        0,
		Limit:         10,
		Keyword:       "希望",
		KeywordEnable: true,
	}
	res, err := elasticsearch.CommentList(ctx, params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func TestCommentCount(t *testing.T) {
	elasticsearch.InitElasticSearch()
	ctx := context.Background()
	params := &model.CommentCountParams{
		BeginTime:      1584288000,
		EndTime:        1584892800,
		Polarity:       "",
		PolarityEnable: false,
		AIDs:           nil,
		APPIDEnable:    false,
		Keyword:        "我们",
		KeywordEnable:  true,
	}
	res, err := elasticsearch.CommentCount(ctx, params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
