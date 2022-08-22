package news_doc_service

import (
	"api-server/pkg/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var esTestDocID = "test-doc-id"

// test doc for: index, search and delete
var esTestDoc = NewsDoc{
	ID:          esTestDocID,
	Link:        "https://test_news_doc_1",
	Title:       "test doc title",
	Description: "test description",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

// test doc for: update
var esUpdateTestDoc = NewsDoc{
	ID:          esTestDocID,
	Title:       "updated doc title",
	Description: "updated description",
}

func TestESDAOIndex(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_index.go")
	fmt.Println("> Index(ctx context.Context, doc *NewsDoc) error")

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
	err = esDAO.Index(ctx, &esTestDoc)
	assert.NoError(t, err)
}

func TestESDAOSearch(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_search.go")
	fmt.Println("> Search(ctx context.Context, query string) (*ESDAOSearchResponse, error)")

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
		`{"from":0, "size":10, "query": {"term": {"id": "`+esTestDocID+`"}}}`,
	)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(resp.Results))

	// check if result correct
	assert.Equal(
		t,
		resp.Results[0].Link,
		esTestDoc.Link,
	)
	assert.Equal(
		t,
		resp.Results[0].Title,
		esTestDoc.Title,
	)
	assert.Equal(
		t,
		resp.Results[0].Description,
		esTestDoc.Description,
	)
	assert.Empty(
		t,
		utils.StringSlicesXOR(resp.Results[0].Authors, esTestDoc.Authors),
	)
	assert.Equal(
		t,
		resp.Results[0].Category,
		esTestDoc.Category,
	)
}

func TestESDAOUpdate(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_update.go")
	fmt.Println("> Update(ctx context.Context, doc *NewsDoc) error")

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

	// start updating
	ctx := context.Background()
	err = esDAO.Update(ctx, &esUpdateTestDoc)
	assert.NoError(t, err)

	// check update by search
	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"id": "`+esTestDocID+`"}}}`,
	)
	assert.NoError(t, err)

	// check if result correct
	assert.Equal(
		t,
		resp.Results[0].Title,
		esUpdateTestDoc.Title,
	)
	assert.Equal(
		t,
		resp.Results[0].Description,
		esUpdateTestDoc.Description,
	)
}

func TestESDAODelete(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_delete.go")
	fmt.Println("> Delete(ctx context.Context, docID string) error")

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

	// start deleting
	ctx := context.Background()
	err = esDAO.Delete(ctx, esTestDocID)
	assert.NoError(t, err)

	// check delete by search
	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"id": "`+esTestDocID+`"}}}`,
	)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp.Results))
}
