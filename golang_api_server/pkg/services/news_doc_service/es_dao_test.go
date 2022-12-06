package news_doc_service_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/services/news_doc_service"
	"api-server/pkg/utils"

	"github.com/stretchr/testify/assert"
)

var esTestDocID = "test-doc-id"

// test doc for: index, search and delete
var esTestDoc = news_doc_service.NewsDoc{
	UUID:        esTestDocID,
	Link:        "https://test_news_doc_1",
	Title:       "test doc title",
	Description: "test description",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

// test doc for: update
var esUpdateTestDoc = news_doc_service.NewsDoc{
	UUID:        esTestDocID,
	Title:       "updated doc title",
	Description: "updated description",
}

func InitESDAO() (news_doc_service.NewDocESDAO, error) {

	// check if env missing
	esURL := os.Getenv("ES_URL")
	if strings.TrimSpace(esURL) == "" {
		return nil, errors.New("missing env ES_URL")
	}

	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	if strings.TrimSpace(esURL) == "" {
		return nil, errors.New("missing env ES_NEWS_DOC_INDEX_ALIAS")
	}

	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")
	if len(esSearchIndex) == 0 {
		return nil, errors.New("missing env ES_NEWS_DOC_SEARCH_ALIAS")
	}

	// es dao
	rawESDAO, err := elasticsearch_data_access.NewESDAO(esURL, esIndexIndex, esSearchIndex)
	if err != nil {
		return nil, err
	}

	esDAO := news_doc_service.NewNewsDocESDAO(*rawESDAO)

	return esDAO, nil
}

func TestESDAOReal(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	t.Run("subtestESDAOIndexReal", subtestESDAOIndexReal)
	t.Run("subtestESDAOSearchReal", subtestESDAOSearchReal)
	t.Run("subTestESDAOUpdateReal", subTestESDAOUpdateReal)
	t.Run("subTestESDAODeleteReal", subTestESDAODeleteReal)
}

func subtestESDAOIndexReal(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_index.go")
	fmt.Println("> Index(ctx context.Context, doc *NewsDoc) error")

	// start indexing
	ctx := context.Background()

	esDAO, err := InitESDAO()
	assert.NoError(t, err)

	err = esDAO.Index(ctx, &esTestDoc)
	assert.NoError(t, err)
}

func subtestESDAOSearchReal(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_search.go")
	fmt.Println("> Search(ctx context.Context, query string) (*ESDAOSearchResponse, error)")

	// run search to get random 10 docs
	ctx := context.Background()

	esDAO, err := InitESDAO()
	assert.NoError(t, err)

	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"uuid": "`+esTestDocID+`"}}}`,
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

func subTestESDAOUpdateReal(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_update.go")
	fmt.Println("> Update(ctx context.Context, doc *NewsDoc) error")

	// start updating
	ctx := context.Background()

	esDAO, err := InitESDAO()
	assert.NoError(t, err)

	err = esDAO.Update(ctx, &esUpdateTestDoc)
	assert.NoError(t, err)

	// check update by search
	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"uuid": "`+esTestDocID+`"}}}`,
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

func subTestESDAODeleteReal(t *testing.T) {

	fmt.Println("Test services/news_doc_service/es_dao_delete.go")
	fmt.Println("> Delete(ctx context.Context, docID string) error")

	// start deleting
	ctx := context.Background()

	esDAO, err := InitESDAO()
	assert.NoError(t, err)

	err = esDAO.Delete(ctx, esTestDocID)
	assert.NoError(t, err)

	// check delete by search
	resp, err := esDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"uuid": "`+esTestDocID+`"}}}`,
	)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp.Results))
}
