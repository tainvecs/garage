package news_doc_service

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
	"context"
)

// api request and response
type SearchRequest struct {
	Query string `json:"query" form:"query"`
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
}

type SearchResponse struct {
	Total int                  `json:"total"`
	Docs  []*SearchResponseDoc `json:"docs"`
}

type SearchResponseDoc struct {
	UUID        string   `json:"uuid"`
	Link        string   `json:"link"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Category    string   `json:"category"`
}

// api handling function
type SearchFunc func(ctx context.Context, request *SearchRequest) (*SearchResponse, error)

func NewSearchFunc(esDAO *ESDAO) SearchFunc {
	return func(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {

		response := SearchResponse{
			Total: 0,
			Docs:  make([]*SearchResponseDoc, 0),
		}

		// build es search query
		esSearchParams := ESSearchParameters{
			Query: request.Query,
			Page:  request.Page,
			Limit: request.Limit,
		}
		esSearchQuery, err := BuildESSearchQuery(&esSearchParams)
		if err != nil {
			return &response, err
		}

		// run es search
		esSearchQueryStr, err := elasticsearch_data_access.ESQueryToString(esSearchQuery)
		if err != nil {
			return &response, err
		}

		searchResp, err := esDAO.Search(ctx, esSearchQueryStr)
		if err != nil {
			return &response, err
		}

		// prepare response
		response.Total = searchResp.Total

		for _, d := range searchResp.Results {
			response.Docs = append(
				response.Docs,
				&SearchResponseDoc{
					UUID:        d.UUID,
					Link:        d.Link,
					Title:       d.Title,
					Description: d.Description,
					Authors:     d.Authors,
					Category:    d.Category,
				},
			)
		}

		return &response, nil
	}
}
