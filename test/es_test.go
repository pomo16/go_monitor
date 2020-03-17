package test

import (
	"context"
	"fmt"
	"gowatcher/go_monitor/service/elasticsearch"
	"testing"
)

func TestES(t *testing.T) {
	elasticsearch.InitElasticSearch()
}

func TestCommentList(t *testing.T) {
	elasticsearch.InitElasticSearch()
	ctx := context.Background()
	res, err := elasticsearch.QueryByMainID(ctx, "27ed3b0a86895e1731d37ddcf5202318")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res[0])
	}
}
