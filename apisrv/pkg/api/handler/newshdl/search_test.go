package newshdl_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

var searchTestDoc = newssvc.ESNewsDoc{
	UUID:        "test-doc-id-1",
	Link:        "https://test_news_doc_1",
	Title:       "test doc title 1",
	Description: "test description 1",
	Authors:     []string{"a1", "a2"},
	Category:    "CATEGORIY",
}

var softDeletedTime = time.Now()
var searchTestDocSoftDeleted = newssvc.ESNewsDoc{
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

	// prerequisite
	esDAO, err := initESDAO()
	assert.NoError(t, err)
	svcESDAO := newssvc.NewESDAO(esDAO)

	ctx := context.Background()
	searchFunc := newshdl.NewSearchFunc(svcESDAO)

	// start test data indexing
	err = svcESDAO.Index(ctx, &searchTestDoc)
	assert.NoError(t, err)
	err = svcESDAO.Index(ctx, &searchTestDocSoftDeleted)
	assert.NoError(t, err)

	// start testing
	request := newshdl.SearchRequest{
		Query: "test description | test title",
		Page:  0,
		Limit: 10,
	}
	resp, err := searchFunc(ctx, &request)
	assert.NoError(t, err)
	assert.Equal(t, resp.Total, 1)
	assert.Equal(t, resp.Docs[0].Title, "test doc title 1")

	// clean up by delete test data
	err = esDAO.Delete(ctx, searchTestDoc.UUID)
	assert.NoError(t, err)
	err = esDAO.Delete(ctx, searchTestDocSoftDeleted.UUID)
	assert.NoError(t, err)
}
