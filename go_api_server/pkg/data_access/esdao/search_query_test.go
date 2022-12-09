package esdao_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/apisrv/pkg/utils/strutils"
)

func TestNewTermQuery(t *testing.T) {

	ansStr := `
        {
            "term":{
                "authors":{
                    "value":"a3",
                    "boost":"1.0"
                }
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es term query
	term := esdao.Term{
		Value: "a3",
		Boost: "1.0",
	}
	termQuery, err := esdao.NewTermQuery("authors", term)
	assert.NoError(t, err)

	// term query -> str
	termStr, err := esdao.QueryToString(termQuery)
	assert.NoError(t, err)
	termStr = strutils.TrimAllIndent(termStr)

	assert.Equal(t, ansStr, termStr)
}

func TestNewTermsQuery(t *testing.T) {

	ansStr := `
        {
            "terms":{
                "authors":[
                    "a2",
                    "a3"
                ],
                "boost":"1.0"
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es terms query
	termsQuery, err := esdao.NewTermsQuery(
		"authors",
		[]string{"a2", "a3"},
		"1.0",
	)
	assert.NoError(t, err)

	// terms query -> str
	termsStr, err := esdao.QueryToString(termsQuery)
	assert.NoError(t, err)
	termsStr = strutils.TrimAllIndent(termsStr)

	assert.Equal(t, ansStr, termsStr)
}

func TestNewMatchQuery(t *testing.T) {

	ansStr := `
        {
            "match":{
                "description":{
                    "query":"testdescription",
                    "boost":"1.0",
                    "operator":"and"
                }
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es match query
	match := esdao.Match{
		Query:    "test description",
		Boost:    "1.0",
		Operator: "and",
	}
	matchQuery, err := esdao.NewMatchQuery("description", match)
	assert.NoError(t, err)

	// match query -> str
	matchStr, err := esdao.QueryToString(matchQuery)
	assert.NoError(t, err)
	matchStr = strutils.TrimAllIndent(matchStr)

	assert.Equal(t, ansStr, matchStr)
}

func TestNewMatchPhraseQuery(t *testing.T) {

	ansStr := `
        {
            "match_phrase":{
                "description":{
                    "query":"test description",
                    "boost":"1.0",
                    "slop":"0"
                }
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es match phrase query
	matchPhrase := esdao.MatchPhrase{
		Query: "test description",
		Boost: "1.0",
		Slop:  "0",
	}
	matchPhraseQuery, err := esdao.NewMatchPhraseQuery("description", matchPhrase)
	assert.NoError(t, err)

	// match phrase query -> str
	matchPhraseStr, err := esdao.QueryToString(matchPhraseQuery)
	assert.NoError(t, err)
	matchPhraseStr = strutils.TrimAllIndent(matchPhraseStr)

	assert.Equal(t, ansStr, matchPhraseStr)
}

func TestNewMultiMatchQuery(t *testing.T) {

	ansStr := `
        {
            "multi_match":{
                "query":"test",
                "type":"best_fields",
                "fields":[
                    "title",
                    "description"
                ]
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es multi match query

	multiMatch := esdao.MultiMatch{
		Query: "test",
		Type:  "best_fields",
		Field: []string{"title", "description"},
	}
	multiMatchQuery, err := esdao.NewMultiMatchQuery(multiMatch)
	assert.NoError(t, err)

	// multi match query -> str
	multiMatchStr, err := esdao.QueryToString(multiMatchQuery)
	assert.NoError(t, err)
	multiMatchStr = strutils.TrimAllIndent(multiMatchStr)

	assert.Equal(t, ansStr, multiMatchStr)
}

func TestNewConstantScoreQuery(t *testing.T) {

	ansStr := `
        {
            "constant_score":{
                "filter":{
                    "term":{
                        "category":{
                            "value":"CATEGORIY",
                            "boost":"1.0"
                        }
                    }
                },
                "boost":"100.0"
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new term query for
	term := esdao.Term{
		Value: "CATEGORIY",
		Boost: "1.0",
	}
	termQuery, err := esdao.NewTermQuery("category", term)
	assert.NoError(t, err)

	// new es constant score query
	constantScore := esdao.ConstantScore{
		Filter: termQuery,
		Boost:  "100.0",
	}
	constantScoreQuery, err := esdao.NewConstantScoreQuery(constantScore)
	assert.NoError(t, err)

	// constant score query -> str
	constantScoreStr, err := esdao.QueryToString(constantScoreQuery)
	assert.NoError(t, err)
	constantScoreStr = strutils.TrimAllIndent(constantScoreStr)

	assert.Equal(t, ansStr, constantScoreStr)
}

func TestNewSimpleQueryStringQuery(t *testing.T) {

	ansStr := `
        {
            "simple_query_string":{
                "query":"test title|test description",
                "fields":[
                    "title^3",
                    "description"
                ],
                "default_operator":"and"
            }
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// new es simple query string query
	sqs := esdao.SimpleQueryString{
		Query:           "test title | test description",
		Fields:          []string{"title^3", "description"},
		DefaultOperator: "and",
	}
	sqsQuery, err := esdao.NewSimpleQueryStringQuery(sqs)
	assert.NoError(t, err)

	// simple query string query -> str
	sqsStr, err := esdao.QueryToString(sqsQuery)
	assert.NoError(t, err)
	sqsStr = strutils.TrimAllIndent(sqsStr)

	assert.Equal(t, ansStr, sqsStr)
}
