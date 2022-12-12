package newssvc

import "github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"

// ESSearchParameters is the search params for building elasticsearch query
type ESSearchParameters struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

// BuildESSearchQuery build the default search query for News Docs
func BuildESSearchQuery(params *ESSearchParameters) (*esdao.QueryBody, error) {

	// bool query
	boolQuery := esdao.BoolQuery{
		Bool: esdao.Bool{
			MustNot:            make([]interface{}, 0),
			Should:             make([]interface{}, 0),
			MinimumShouldMatch: "1",
		},
	}

	// filter soft deleted docs
	exists := esdao.Exists{
		Field: "deleted_at",
	}
	existsQuery, err := esdao.NewExistsQuery(exists)
	if err != nil {
		return nil, err
	}
	boolQuery.Bool.MustNot = append(boolQuery.Bool.MustNot, existsQuery)

	// simple query string query
	simpleQueryString := esdao.SimpleQueryString{
		Query: params.Query,
		Fields: []string{
			"title^3.0",
			"description^1.0",
		},
		DefaultOperator: "and",
	}
	simpleQueryStringQuery, err := esdao.NewSimpleQueryStringQuery(simpleQueryString)
	if err != nil {
		return nil, err
	}
	boolQuery.Bool.Should = append(boolQuery.Bool.Should, simpleQueryStringQuery)

	// result query
	esQueryBody := esdao.QueryBody{
		From:     params.Page * params.Limit,
		Size:     params.Limit,
		MinScore: 0.00001,
		Query:    boolQuery,
	}

	return &esQueryBody, nil
}
