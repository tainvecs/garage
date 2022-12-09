package esdao_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/utils/strutils"
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
	ansStr = strutils.TrimAllIndent(ansStr)

	respStr, err := esdao.QueryToString(esdao.RawSearchResponse{})
	assert.NoError(t, err)
	respStr = strutils.TrimAllIndent(respStr)

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
	ansStr = strutils.TrimAllIndent(ansStr)

	hitStr, err := esdao.QueryToString(esdao.RawSearchHit{})
	assert.NoError(t, err)
	hitStr = strutils.TrimAllIndent(hitStr)

	assert.Equal(t, ansStr, hitStr)
}
