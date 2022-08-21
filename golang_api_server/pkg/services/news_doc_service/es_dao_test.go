package news_doc_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestESDAOIndex(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_index.go")
	fmt.Println("> Index(ctx context.Context, doc *NewsDoc)")

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
	esDAO, err := NewESDAO(esURL, esIndexIndex, esSearchIndex)
	assert.NoError(t, err)

	// start indexing
	ctx := context.Background()
	timeNow := time.Now()

	doc := NewsDoc{
		ID:          "test-doc-id",
		Link:        "https://test_news_doc_1",
		Title:       "test doc title",
		Description: "test description",
		Authors:     []string{"a1", "a2"},
		Category:    "test_category",
		CreatedAt:   &timeNow,
	}
	err = esDAO.Index(ctx, &doc)
	assert.NoError(t, err)
}

func TestESDAOSearch(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_search.go")
	fmt.Println("> Search(ctx context.Context, query string)")

	// check if env missing
	esURL := os.Getenv("ES_URL")
	if strings.TrimSpace(esURL) == "" {
		panic(errors.New("missing env ES_URL"))
	}

	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	assert.NotEmpty(t, esIndexIndex)

	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")
	if strings.TrimSpace(esURL) == "" {
		panic(errors.New("missing env ES_NEWS_DOC_SEARCH_ALIAS"))
	}

	// prerequisite
	esDAO, err := NewESDAO(esURL, esIndexIndex, esSearchIndex)
	assert.NoError(t, err)

	// run search to get random 10 docs
	ctx := context.Background()
	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"match_all": {}}}`,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	mar, err := json.MarshalIndent(resp, "", "\t")
	assert.NoError(t, err)
	fmt.Println(string(mar))
}

func TestESDAOUpdate(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_update.go")
	fmt.Println("> Update(ctx context.Context, doc *NewsDoc)")

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
	esDAO, err := NewESDAO(esURL, esIndexIndex, esSearchIndex)
	assert.NoError(t, err)

	// start indexing
	ctx := context.Background()

	doc := NewsDoc{
		ID:          "test-doc-id",
		Title:       "updated doc title",
		Description: "updated description",
	}
	err = esDAO.Update(ctx, &doc)
	assert.NoError(t, err)
}

func TestESDAODelete(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_delete.go")
	fmt.Println("> Delete(ctx context.Context, docID string)")

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
	esDAO, err := NewESDAO(esURL, esIndexIndex, esSearchIndex)
	assert.NoError(t, err)

	// start indexing
	ctx := context.Background()
	err = esDAO.Delete(ctx, "test-doc-id")
	assert.NoError(t, err)
}
