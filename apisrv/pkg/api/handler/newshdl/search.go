package newshdl

import (
	"context"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

// SearchRequest is the struct of api struct request
type SearchRequest struct {
	Query string `json:"query" form:"query" binding:"required"`
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
}

// SearchResponse is the struct of api struct response
type SearchResponse struct {
	Total int                  `json:"total"`
	Docs  []*SearchResponseDoc `json:"docs"`
}

// SearchResponseDoc is part of SearchResponse
type SearchResponseDoc struct {
	Link        string   `json:"link"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Category    string   `json:"category"`
}

// SearchFunc is the api search request handling function
type SearchFunc func(ctx context.Context, request *SearchRequest) (*SearchResponse, error)

// NewSearchFunc instantiate a SearchFunc for handler
func NewSearchFunc(esDAO newssvc.ESDAO) SearchFunc {
	return func(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {

		response := SearchResponse{
			Total: 0,
			Docs:  make([]*SearchResponseDoc, 0),
		}

		// build es search query
		params := newssvc.ESSearchParameters{
			Query: request.Query,
			Page:  request.Page,
			Limit: request.Limit,
		}
		query, err := newssvc.BuildESSearchQuery(&params)
		if err != nil {
			return &response, err
		}

		// run es search
		queryStr, err := esdao.QueryToString(query)
		if err != nil {
			return &response, err
		}

		searchResp, err := esDAO.Search(ctx, queryStr)
		if err != nil {
			return &response, err
		}

		// prepare response
		response.Total = searchResp.Total

		for _, d := range searchResp.Results {
			response.Docs = append(
				response.Docs,
				&SearchResponseDoc{
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
