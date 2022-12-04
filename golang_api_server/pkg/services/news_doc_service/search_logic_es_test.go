package news_doc_service_test

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/services/news_doc_service"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var esSearchTestDoc = news_doc_service.NewsDoc{
	UUID:        "test-doc-id",
	Link:        "https://test_news_doc_1",
	Title:       "test doc title 1",
	Description: "test description 1",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

var softDeletedTime = time.Now()
var esSearchTestDocSoftDeleted = news_doc_service.NewsDoc{
	UUID:        "test-doc-id",
	Link:        "https://test_news_doc_2",
	Title:       "test doc title 2",
	Description: "test description 2",
	Authors:     []string{"a3", "a4"},
	Category:    "CATEGORIY",
	DeletedAt:   &softDeletedTime,
}

func TestBuildESSearchQuery(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	fmt.Println("Test services/news_doc_service/search_logic_es.go")
	fmt.Println("> BuildESSearchQuery(params *ESSearchParameters) (*elasticsearch_data_access.QueryBody, error)")

	// check if env missing
	esURL := os.Getenv("ES_URL")
	if strings.TrimSpace(esURL) == "" {
		panic(errors.New("missing env ES_URL"))
	}

	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	if strings.TrimSpace(esURL) == "" {
		panic(errors.New("missing env ES_NEWS_DOC_INDEX_ALIAS"))
	}

	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")
	assert.NotEmpty(t, esSearchIndex)

	// prerequisite
	esDAO, err := news_doc_service.NewESDAO(esURL, esIndexIndex, esSearchIndex)
	assert.NoError(t, err)

	// start indexing
	ctx := context.Background()

	err = esDAO.Index(ctx, &esSearchTestDoc)
	assert.NoError(t, err)

	err = esDAO.Index(ctx, &esSearchTestDocSoftDeleted)
	assert.NoError(t, err)

	// start search test
	params := news_doc_service.ESSearchParameters{
		Query: "test description | test title",
		Page:  0,
		Limit: 10,
	}
	searchQuery, err := news_doc_service.BuildESSearchQuery(&params)
	assert.NoError(t, err)

	searchQueryStr, err := elasticsearch_data_access.ESQueryToString(searchQuery)
	assert.NoError(t, err)

	resp, err := esDAO.Search(ctx, searchQueryStr)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(resp.Results))

	// check if result correct
	assert.Equal(
		t,
		resp.Results[0].UUID,
		esSearchTestDocSoftDeleted.UUID,
	)
}
