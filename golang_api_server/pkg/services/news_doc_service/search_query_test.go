package news_doc_service_test

import (
	"fmt"
	"testing"

	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/services/news_doc_service"
	"api-server/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestBuildESSearchQuery(t *testing.T) {

	fmt.Println("Test services/news_doc_service/search_query.go")
	fmt.Println("> BuildESSearchQuery(params *ESSearchParameters) (*elasticsearch_data_access.QueryBody, error)")

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
	params := news_doc_service.ESSearchParameters{
		Query: "test description | test title",
		Page:  0,
		Limit: 10,
	}
	searchQuery, err := news_doc_service.BuildESSearchQuery(&params)
	assert.NoError(t, err)

	// check ans
	searchQueryStr, err := elasticsearch_data_access.ESQueryToString(searchQuery)
	assert.NoError(t, err)
	searchQueryStr, err = utils.StringTrimAllIndent(searchQueryStr)
	assert.NoError(t, err)
	ansStr, err = utils.StringTrimAllIndent(ansStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, searchQueryStr)
}
