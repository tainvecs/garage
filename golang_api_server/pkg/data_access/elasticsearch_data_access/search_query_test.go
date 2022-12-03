package elasticsearch_data_access_test

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTermQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewTermQuery(field string, term Term) (*TermQuery, error)")

	ansStr := `{"term":{"authors":{"value":"a3","boost":"1.0"}}}`

	// new es term query
	term := elasticsearch_data_access.Term{
		Value: "a3",
		Boost: "1.0",
	}
	termQuery, err := elasticsearch_data_access.NewTermQuery("authors", term)
	assert.NoError(t, err)

	// term query -> str
	termStr, err := elasticsearch_data_access.ESQueryToString(termQuery)
	assert.NoError(t, err)
	termStr, err = utils.StringTrimAllIndent(termStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, termStr)
}

func TestNewTermsQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewTermsQuery(field string, values []string, boost string) (*TermsQuery, error)")

	ansStr := `{"terms":{"authors":["a2","a3"],"boost":"1.0"}}`

	// new es terms query
	termsQuery, err := elasticsearch_data_access.NewTermsQuery(
		"authors",
		[]string{"a2", "a3"},
		"1.0",
	)
	assert.NoError(t, err)

	// terms query -> str
	termsStr, err := elasticsearch_data_access.ESQueryToString(termsQuery)
	assert.NoError(t, err)
	termsStr, err = utils.StringTrimAllIndent(termsStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, termsStr)
}

func TestNewMatchQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewMatchQuery(field string, match Match) (*MatchQuery, error)")

	ansStr := `{"match":{"description":{"query":"testdescription","boost":"1.0","operator":"and"}}}`

	// new es match query
	match := elasticsearch_data_access.Match{
		Query:    "test description",
		Boost:    "1.0",
		Operator: "and",
	}
	matchQuery, err := elasticsearch_data_access.NewMatchQuery("description", match)
	assert.NoError(t, err)

	// match query -> str
	matchStr, err := elasticsearch_data_access.ESQueryToString(matchQuery)
	assert.NoError(t, err)
	matchStr, err = utils.StringTrimAllIndent(matchStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, matchStr)
}

func TestNewMatchPhraseQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewMatchPhraseQuery(field string, matchPhrase MatchPhrase) (*MatchPhraseQuery, error)")

	ansStr := `{"match_phrase":{"description":{"query":"testdescription","boost":"1.0","slop":"0"}}}`

	// new es match phrase query
	matchPhrase := elasticsearch_data_access.MatchPhrase{
		Query: "test description",
		Boost: "1.0",
		Slop:  "0",
	}
	matchPhraseQuery, err := elasticsearch_data_access.NewMatchPhraseQuery("description", matchPhrase)
	assert.NoError(t, err)

	// match phrase query -> str
	matchPhraseStr, err := elasticsearch_data_access.ESQueryToString(matchPhraseQuery)
	assert.NoError(t, err)
	matchPhraseStr, err = utils.StringTrimAllIndent(matchPhraseStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, matchPhraseStr)
}

func TestNewMultiMatchQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewMultiMatchQuery(multiMatch MultiMatch) (*MultiMatchQuery, error)")

	ansStr := `{"multi_match":{"query":"test","type":"best_fields","fields":["title","description"]}}`

	// new es multi match query

	multiMatch := elasticsearch_data_access.MultiMatch{
		Query: "test",
		Type:  "best_fields",
		Field: []string{"title", "description"},
	}
	multiMatchQuery, err := elasticsearch_data_access.NewMultiMatchQuery(multiMatch)
	assert.NoError(t, err)

	// multi match query -> str
	multiMatchStr, err := elasticsearch_data_access.ESQueryToString(multiMatchQuery)
	assert.NoError(t, err)
	multiMatchStr, err = utils.StringTrimAllIndent(multiMatchStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, multiMatchStr)
}

func TestNewConstantScoreQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewConstantScoreQuery(constantScore ConstantScore) (*ConstantScoreQuery, error)")

	ansStr := `{"constant_score":{"filter":{"term":{"category":{"value":"CATEGORIY","boost":"1.0"}}},"boost":"100.0"}}`

	// new term query for
	term := elasticsearch_data_access.Term{
		Value: "CATEGORIY",
		Boost: "1.0",
	}
	termQuery, err := elasticsearch_data_access.NewTermQuery("category", term)
	assert.NoError(t, err)

	// new es constant score query
	constantScore := elasticsearch_data_access.ConstantScore{
		Filter: termQuery,
		Boost:  "100.0",
	}
	constantScoreQuery, err := elasticsearch_data_access.NewConstantScoreQuery(constantScore)
	assert.NoError(t, err)

	// constant score query -> str
	constantScoreStr, err := elasticsearch_data_access.ESQueryToString(constantScoreQuery)
	assert.NoError(t, err)
	constantScoreStr, err = utils.StringTrimAllIndent(constantScoreStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, constantScoreStr)
}

func TestNewSimpleQueryStringQuery(t *testing.T) {

	fmt.Println("Test data_access/elasticsearch_data_access/search_query.go")
	fmt.Println("> NewSimpleQueryStringQuery(simpleQueryString SimpleQueryString) (*SimpleQueryStringQuery, error)")

	ansStr := `{"simple_query_string":{"query":"testtitle|testdescription","fields":["title^3","description"],"default_operator":"and"}}`

	// new es simple query string query
	sqs := elasticsearch_data_access.SimpleQueryString{
		Query:           "test title | test description",
		Fields:          []string{"title^3", "description"},
		DefaultOperator: "and",
	}
	sqsQuery, err := elasticsearch_data_access.NewSimpleQueryStringQuery(sqs)
	assert.NoError(t, err)

	// simple query string query -> str
	sqsStr, err := elasticsearch_data_access.ESQueryToString(sqsQuery)
	assert.NoError(t, err)
	sqsStr, err = utils.StringTrimAllIndent(sqsStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, sqsStr)
}
