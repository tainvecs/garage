package news_doc_handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"api-server/pkg/handler/news_doc_handler"
	"api-server/pkg/services/news_doc_service"

	"github.com/stretchr/testify/assert"
)

var esSearchTestDoc = news_doc_service.NewsDoc{
	UUID:        "test-doc-id-1",
	Link:        "https://test_news_doc_1",
	Title:       "test doc title 1",
	Description: "test description 1",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

var softDeletedTime = time.Now()
var esSearchTestDocSoftDeleted = news_doc_service.NewsDoc{
	UUID:        "test-doc-id-2",
	Link:        "https://test_news_doc_2",
	Title:       "test doc title 2",
	Description: "test description 2",
	Authors:     []string{"a3", "a4"},
	Category:    "CATEGORIY",
	DeletedAt:   &softDeletedTime,
}

func TestNewSearchFunc(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	fmt.Println("Test api/news_doc_handler/search.go")
	fmt.Println("> NewSearchFunc(esDAO *ESDAO) SearchFunc func(ctx context.Context, request *SearchRequest) (*SearchResponse, error)")

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

	searchFunc := news_doc_handler.NewSearchFunc(esDAO)

	// start test data indexing
	ctx := context.Background()

	err = esDAO.Index(ctx, &esSearchTestDoc)
	assert.NoError(t, err)
	err = esDAO.Index(ctx, &esSearchTestDocSoftDeleted)
	assert.NoError(t, err)

	// start testing
	request := news_doc_handler.SearchRequest{
		Query: "test description | test title",
		Page:  0,
		Limit: 10,
	}
	resp, err := searchFunc(ctx, &request)
	assert.NoError(t, err)

	// clean up by delete test data
	err = esDAO.Delete(ctx, esSearchTestDoc.UUID)
	assert.NoError(t, err)
	err = esDAO.Delete(ctx, esSearchTestDocSoftDeleted.UUID)
	assert.NoError(t, err)

	// print debug
	mar, err := json.MarshalIndent(resp, "", "\t")
	assert.NoError(t, err)
	fmt.Println(string(mar))
}
