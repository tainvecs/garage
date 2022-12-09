package esdao_test

import (
	"apisrv/pkg/data_access/esdao"
	"apisrv/pkg/utils/strutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawSearchResponse(t *testing.T) {

	// test search response marshal result
	ansStr := `
        {
            "took":0,
            "timed_out":false,
            "Shards":{
                "total":0,
                "successful":0,
                "skipped":0,
                "failed":0
            },
            "hits":{
                "total":{
                    "value":0,
                    "relation":""
                },
                "mas_score":0,
                "hits":null
            }
        }
    `
	ansStr, err := strutils.TrimAllIndent(ansStr)
	assert.NoError(t, err)

	respStr, err := esdao.QueryToString(esdao.RawSearchResponse{})
	assert.NoError(t, err)
	respStr, err = strutils.TrimAllIndent(respStr)
	assert.NoError(t, err)

	assert.Equal(t, ansStr, respStr)
}

func TestRawSearchHit(t *testing.T) {

	// test search hit marshal result
	ansStr := `
        {
            "_index":"",
            "_type":"",
            "_id":"",
            "_score":0
        }
    `
	ansStr, err := strutils.TrimAllIndent(ansStr)
	assert.NoError(t, err)

	hitStr, err := esdao.QueryToString(esdao.RawSearchHit{})
	assert.NoError(t, err)
	hitStr, err = strutils.TrimAllIndent(hitStr)
	assert.NoError(t, err)

	assert.Equal(t, ansStr, hitStr)
}
