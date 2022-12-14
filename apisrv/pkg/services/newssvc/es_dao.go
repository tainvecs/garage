package newssvc

import (
	"context"
	"time"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
)

// ESNewsDoc is the struct of news docs stored in elasticsearch
type ESNewsDoc struct {
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

// ESDAO is the elasticsearch data access object for news docs
type ESDAO interface {
	Index(ctx context.Context, doc *ESNewsDoc) error
	Search(ctx context.Context, query string) (*ESSearchResponse, error)
	Update(ctx context.Context, doc *ESNewsDoc) error
	Delete(ctx context.Context, docID string) error
}

// esDAO use the esdao.DataAccessObject to access elasticsearch
type esDAO struct {
	DataAccessObject *esdao.DataAccessObject
}

// NewESDAO instansite a new ESDAO
func NewESDAO(dao *esdao.DataAccessObject) ESDAO {
	return &esDAO{DataAccessObject: dao}
}
