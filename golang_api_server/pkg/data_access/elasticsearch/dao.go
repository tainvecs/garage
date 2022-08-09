package es_dao

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESDAO interface {
	Index(ctx context.Context, doc *Document) error
	Search(ctx context.Context, query string) (*DAOSearchResponse, error)
	Update(ctx context.Context, doc *Document) error
	Delete(ctx context.Context, docID string) error
}

type DAO struct {
	Client      *ESClient
	IndexIndex  string
	SearchIndex string
}

type DAOSearchResponse struct {
	Total   int
	Results []*Document
	Scores  []float32
}

func NewDAO(esURL, esIndexIndex, esSearchIndex string) (ESDAO, error) {

	// new elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// new DAO
	dao := DAO{
		Client:      &ESClient{Client: client},
		IndexIndex:  esIndexIndex,
		SearchIndex: esSearchIndex,
	}

	return &dao, nil
}
