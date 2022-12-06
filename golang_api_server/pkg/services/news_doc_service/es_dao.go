package news_doc_service

import (
	"context"
	"time"

	es_data_access "api-server/pkg/data_access/elasticsearch_data_access"
)

// data element
type NewsDoc struct {
	UUID        string     `json:"uuid"`
	Link        string     `json:"link,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Authors     []string   `json:"authors,omitempty"`
	Category    string     `json:"category,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// ES data access object interface
type NewDocESDAO interface {
	Index(ctx context.Context, doc *NewsDoc) error
	Search(ctx context.Context, query string) (*ESDAOSearchResponse, error)
	Update(ctx context.Context, doc *NewsDoc) error
	Delete(ctx context.Context, docID string) error
}

// ES data access object
type newsDocESDAO struct {
	Client      *es_data_access.ESClient
	IndexIndex  string
	SearchIndex string
}

func NewNewsDocESDAO(esDAO es_data_access.ESDAO) NewDocESDAO {
	return &newsDocESDAO{
		Client:      esDAO.Client,
		IndexIndex:  esDAO.IndexIndex,
		SearchIndex: esDAO.SearchIndex,
	}
}
