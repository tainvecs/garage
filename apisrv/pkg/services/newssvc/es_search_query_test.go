package newssvc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
	"github.com/tainvecs/garage/apisrv/pkg/utils/strutils"
)

func TestBuildESSearchQuery(t *testing.T) {

	ansStr := `
        {
            "size":10,
            "min_score":0.00001,
            "_source":{

            },
            "query":{
                "bool":{
                    "must_not":[
                        {
                            "exists":{
                                "field":"deleted_at"
                            }
                        }
                    ],
                    "should":[
                        {
                            "simple_query_string":{
                                "query":"testdescription|testtitle",
                                "fields":[
                                    "title^3.0",
                                    "description^1.0"
                                ],
                                "default_operator":"and"
                            }
                        }
                    ],
                    "minimum_should_match":"1"
                }
            },
            "highlight":{

            }
        }
        `

	// build es search query
	params := newssvc.ESSearchParameters{
		Query: "test description | test title",
		Page:  0,
		Limit: 10,
	}
	searchQuery, err := newssvc.BuildESSearchQuery(&params)
	assert.NoError(t, err)

	// check ans
	searchQueryStr, err := esdao.QueryToString(searchQuery)
	assert.NoError(t, err)
	searchQueryStr = strutils.TrimAllIndent(searchQueryStr)
	ansStr = strutils.TrimAllIndent(ansStr)
	assert.Equal(t, ansStr, searchQueryStr)
}
