package elasticsearch_data_access

import (
	"github.com/elastic/go-elasticsearch/v8"
)

// ES data access object
type ESDAO struct {
	Client      *ESClient
	IndexIndex  string
	SearchIndex string
}

func NewESDAO(esURL, esIndexIndex, esSearchIndex string) (*ESDAO, error) {

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
		Client:      &ESClient{Client: client},
		IndexIndex:  esIndexIndex,
		SearchIndex: esSearchIndex,
	}

	return &dao, nil
}
