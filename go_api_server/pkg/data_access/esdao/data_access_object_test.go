package esdao_test

import (
	"apisrv/pkg/data_access/esdao"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testDoc struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func TestDataAccessObject(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	ctx := context.Background()

	// test doc
	docID := "test-doc-id"
	doc := testDoc{ // for: index, search and delete
		ID:    docID,
		Title: "test doc title",
	}
	updateDoc := testDoc{ // for: update
		Title: "updated doc title",
	}

	// check if env missing
	url := os.Getenv("ES_URL")
	assert.NotEmpty(t, strings.TrimSpace(url))

	indexIndex := os.Getenv("ES_UNIT_TEST_INDEX_ALIAS")
	assert.NotEmpty(t, strings.TrimSpace(indexIndex))

	searchIndex := os.Getenv("ES_UNIT_TEST_SEARCH_ALIAS")
	assert.NotEmpty(t, strings.TrimSpace(searchIndex))

	// test new data access object
	dao, err := esdao.New(url, indexIndex, searchIndex)
	assert.NoError(t, err)

	// test index
	err = dao.Index(ctx, docID, doc)
	assert.NoError(t, err)

	// test search
	searchQuery := `{"query": {"term": {"id": "` + docID + `"}}}`
	searchResponse, err := dao.Search(ctx, searchQuery)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(searchResponse.Hits.Hits))

	// test update
	err = dao.Update(ctx, docID, updateDoc)
	assert.NoError(t, err)

	// test delete
	err = dao.Delete(ctx, docID)
	assert.NoError(t, err)
}
