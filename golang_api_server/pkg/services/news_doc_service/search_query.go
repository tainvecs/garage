package news_doc_service

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
)

// es search params and query
type ESSearchParameters struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func BuildESSearchQuery(params *ESSearchParameters) (*elasticsearch_data_access.QueryBody, error) {

	// bool query
	boolQuery := elasticsearch_data_access.BoolQuery{
		Bool: elasticsearch_data_access.Bool{
			MustNot:            make([]interface{}, 0),
			Should:             make([]interface{}, 0),
			MinimumShouldMatch: "1",
		},
	}

	// filter soft deleted docs
	exists := elasticsearch_data_access.Exists{
		Field: "deleted_at",
	}
	existsQuery, err := elasticsearch_data_access.NewExistsQuery(exists)
	if err != nil {
		return nil, err
	}
	boolQuery.Bool.MustNot = append(boolQuery.Bool.MustNot, existsQuery)

	// simple query string query
	simpleQueryString := elasticsearch_data_access.SimpleQueryString{
		Query: params.Query,
		Fields: []string{
			"title^3.0",
			"description^1.0",
		},
		DefaultOperator: "and",
	}
	simpleQueryStringQuery, err := elasticsearch_data_access.NewSimpleQueryStringQuery(simpleQueryString)
	if err != nil {
		return nil, err
	}
	boolQuery.Bool.Should = append(boolQuery.Bool.Should, simpleQueryStringQuery)

	// result query
	esQueryBody := elasticsearch_data_access.QueryBody{
		From:     params.Page * params.Limit,
		Size:     params.Limit,
		MinScore: 0.00001,
		Query:    boolQuery,
	}

	return &esQueryBody, nil
}
