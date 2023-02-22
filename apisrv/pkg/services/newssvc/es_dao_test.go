package newssvc_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
	"github.com/tainvecs/garage/apisrv/pkg/utils/strutils"
)

var testESDAO newssvc.ESDAO

var esTestDocID = "test-doc-id"

// test doc for: index, search and delete
var esTestDoc = newssvc.ESNewsDoc{
	UUID:        esTestDocID,
	Link:        "https://test_news_doc_1",
	Title:       "test doc title",
	Description: "test description",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

// test doc for: update
var esUpdateTestDoc = newssvc.ESNewsDoc{
	UUID:        esTestDocID,
	Title:       "updated doc title",
	Description: "updated description",
}

func initES() {

	// check if env missing
	url := os.Getenv("ES_URL")
	if len(strings.TrimSpace(url)) == 0 {
		panic(errors.New("missing env ES_URL"))
	}

	indexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	if len(strings.TrimSpace(indexIndex)) == 0 {
		panic(errors.New("missing env ES_NEWS_DOC_INDEX_ALIAS"))
	}

	searchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")
	if len(strings.TrimSpace(searchIndex)) == 0 {
		panic(errors.New("missing env ES_NEWS_DOC_SEARCH_ALIAS"))
	}

	// es dao
	dao, err := esdao.New(url, indexIndex, searchIndex)
	if err != nil {
		panic(err)
	}
	testESDAO = newssvc.NewESDAO(dao)
}

func TestESDAOReal(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	initES()

	t.Run("subtestESDAOIndexReal", subtestESDAOIndexReal)
	t.Run("subtestESDAOSearchReal", subtestESDAOSearchReal)
	t.Run("subTestESDAOUpdateReal", subTestESDAOUpdateReal)
	t.Run("subTestESDAODeleteReal", subTestESDAODeleteReal)
}

func subtestESDAOIndexReal(t *testing.T) {

	err := testESDAO.Index(context.Background(), &esTestDoc)
	assert.NoError(t, err)
}

func subtestESDAOSearchReal(t *testing.T) {

	// run search to get random 10 docs
	ctx := context.Background()

	resp, err := testESDAO.Search(
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
		strutils.StringSlicesXOR(resp.Results[0].Authors, esTestDoc.Authors),
	)
	assert.Equal(
		t,
		resp.Results[0].Category,
		esTestDoc.Category,
	)
}

func subTestESDAOUpdateReal(t *testing.T) {

	// start updating
	ctx := context.Background()

	err := testESDAO.Update(ctx, &esUpdateTestDoc)
	assert.NoError(t, err)

	// check update by search
	resp, err := testESDAO.Search(
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

	// start deleting
	ctx := context.Background()

	err := testESDAO.Delete(ctx, esTestDocID)
	assert.NoError(t, err)

	// check delete by search
	resp, err := testESDAO.Search(
		ctx,
		`{"from":0, "size":10, "query": {"term": {"uuid": "`+esTestDocID+`"}}}`,
	)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp.Results))
}
