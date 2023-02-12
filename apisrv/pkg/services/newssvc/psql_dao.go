package newssvc

import (
	"context"
	"time"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
	"gorm.io/gorm"
)

// PsqlNewsDoc is the struct of news docs stored in PostgreSQL database
type PsqlNewsDoc struct {
	ID          int           `gorm:"column:id;primarykey" json:"id"`
	UUID        string        `gorm:"column:uuid" json:"uuid"`
	Link        string        `gorm:"column:link" json:"link,omitempty"`
	Title       string        `gorm:"column:title" json:"title,omitempty"`
	Description string        `gorm:"column:description" json:"description,omitempty"`
	CreatedAt   *time.Time    `gorm:"column:created_at" json:"created_at,omitempty"`
	Category    string        `gorm:"column:category" json:"category,omitempty"`
	Authors     []*PsqlAuthor `gorm:"many2many:news_authors;joinForeignKey:news_id;joinReferences:authors_id" json:"authors,omitempty"`
}

func (n *PsqlNewsDoc) TableName() string {
	return "news"
}

// PsqlAuthor is the struct of authors of news docs
type PsqlAuthor struct {
	ID   int    `gorm:"column:id;primarykey" json:"id"`
	Name string `gorm:"column:name" json:"name,omitempty"`
}

func (a *PsqlAuthor) TableName() string {
	return "authors"
}

// PsqlNewsAuthors is the reference table for News and Authors
type PsqlNewsAuthors struct {
	ID         int `gorm:"column:id;primarykey" json:"id"`
	NewsID     int `gorm:"column:news_id" json:"news_id,omitempty"`
	Authors_ID int `gorm:"column:authors_id" json:"authors_id,omitempty"`
}

func (na *PsqlNewsAuthors) TableName() string {
	return "news_authors"
}

// for the PreloadAssociations in sqldao QueryConfig
const (
	PsqlNewsAuthorsAssociation = "Authors"
)

// PsqlDAO is the psql data access object for news docs
type PsqlDAO interface {
	GetAll(ctx context.Context, queryConf *sqldao.QueryConfig) ([]*PsqlNewsDoc, error)
}

// psqlDAO use the gorm.DB to access sql database
type psqlDAO struct {
	Client *gorm.DB
}

// NewPsqlDAO instansite a new PsqlDAO
func NewPsqlDAO(client *gorm.DB) PsqlDAO {
	return &psqlDAO{Client: client}
}

// GetAll is the service func for getting all news doc from psql
func (dao *psqlDAO) GetAll(ctx context.Context, queryConf *sqldao.QueryConfig) ([]*PsqlNewsDoc, error) {

	var docSlice []*PsqlNewsDoc

	err := queryConf.
		Apply(dao.Client).
		WithContext(ctx).
		Model(PsqlNewsDoc{}).
		Find(&docSlice).Error

	return docSlice, err
}
