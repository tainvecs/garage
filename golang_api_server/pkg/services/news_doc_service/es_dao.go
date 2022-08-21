package news_doc_service

import (
	"context"

	es_data_access "api-server/pkg/data_access/elasticsearch_data_access"

	"github.com/elastic/go-elasticsearch/v8"
)

// ES data access object interface
type esDAO interface {
	Index(ctx context.Context, doc *NewsDoc) error
	Search(ctx context.Context, query string) (*ESDAOSearchResponse, error)
	Update(ctx context.Context, doc *NewsDoc) error
	Delete(ctx context.Context, docID string) error
}

// ES data access object
type ESDAO struct {
	Client      *es_data_access.ESClient
	IndexIndex  string
	SearchIndex string
}

func NewESDAO(esURL, esIndexIndex, esSearchIndex string) (esDAO, error) {

	// new elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// new DAO
	dao := ESDAO{
		Client:      &es_data_access.ESClient{Client: client},
		IndexIndex:  esIndexIndex,
		SearchIndex: esSearchIndex,
	}

	return &dao, nil
}
