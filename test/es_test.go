package test

import (
	"gowatcher/go_monitor/service/elasticsearch"
	"testing"
)

func TestES(t *testing.T) {
	elasticsearch.InitElasticSearch()
}
