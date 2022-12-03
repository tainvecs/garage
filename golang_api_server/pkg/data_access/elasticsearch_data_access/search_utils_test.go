package elasticsearch_data_access_test

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestESQueryToString(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_utils.go")
	fmt.Println("> ESQueryToString(inQuery interface{}) (string, error)")

	ansStr := `{"filter":[{"match_phrase":{"testfield":{"query":"testquery","boost":"1.0","slop":"1"}}}],"minimum_should_match":"1.0","boost":"1.0"}`

	// test es query
	matchPhrase := elasticsearch_data_access.MatchPhrase{
		Query: "test query",
		Boost: "1.0",
		Slop:  "1",
	}
	matchPhraseQuery, err := elasticsearch_data_access.NewMatchPhraseQuery(
		"test field",
		matchPhrase,
	)
	assert.NoError(t, err)

	testESQuery := elasticsearch_data_access.Bool{
		Filter:             make([]interface{}, 0),
		MinimumShouldMatch: "1.0",
		Boost:              "1.0",
	}
	testESQuery.Filter = append(testESQuery.Filter, matchPhraseQuery)

	// test es query -> str
	testESQueryStr, err := elasticsearch_data_access.ESQueryToString(testESQuery)
	assert.NoError(t, err)
	testESQueryStr, err = utils.StringTrimAllIndent(testESQueryStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, testESQueryStr)
}
