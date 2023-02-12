package newshdl

import (
	"context"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

// SearchRequest is the struct of api get request
type GetRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

// SearchResponse is the struct of api get response
type GetResponse struct {
	Total int                    `json:"total"`
	Docs  []*newssvc.PsqlNewsDoc `json:"docs"`
}

// GetFunc is the api get request handling function
type GetFunc func(ctx context.Context, request *GetRequest) (*GetResponse, error)

// NewGetFunc instantiate a GetFunc for handler
func NewGetFunc(psqlDAO newssvc.PsqlDAO) GetFunc {
	return func(ctx context.Context, request *GetRequest) (*GetResponse, error) {

		// response
		response := GetResponse{
			Total: 0,
			Docs:  make([]*newssvc.PsqlNewsDoc, 0),
		}

		// init psql query configs
		conf := sqldao.QueryConfig{
			Offset:              request.Page * request.Limit,
			Limit:               request.Limit,
			PreloadAssociations: []string{},
		}

		// run get
		docs, err := psqlDAO.Get(ctx, &conf)
		if err != nil {
			return &response, err
		}

		// result
		response.Total = len(docs)
		response.Docs = docs

		return &response, nil
	}
}
