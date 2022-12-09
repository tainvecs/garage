package esdao_test

import (
	"apisrv/pkg/data_access/esdao"
	"apisrv/pkg/utils/strutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryToString(t *testing.T) {

	ansStr := `
        {
            "filter":[
                {
                    "match_phrase":{
                        "test_field":{
                            "query":"test query"
                        }
                    }
                }
            ]
        }
    `
	ansStr = strutils.TrimAllIndent(ansStr)

	// test es query -> str
	testESQuery := map[string]interface{}{
		"filter": []map[string]interface{}{
			{
				"match_phrase": map[string]interface{}{
					"test_field": map[string]string{
						"query": "test query",
					},
				},
			},
		},
	}
	testESQueryStr, err := esdao.QueryToString(testESQuery)
	assert.NoError(t, err)
	testESQueryStr = strutils.TrimAllIndent(testESQueryStr)

	assert.Equal(t, ansStr, testESQueryStr)
}
