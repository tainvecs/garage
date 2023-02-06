package newssvc

import "time"

// PsqlNewsDoc is the struct of news docs stored in PostgreSQL database
type PsqlNewsDoc struct {
	ID          int           `gorm:"column:id;primarykey"`
	UUID        string        `gorm:"column:uuid"`
	Link        string        `gorm:"column:link"`
	Title       string        `gorm:"column:title"`
	Description string        `gorm:"column:description"`
	CreatedAt   *time.Time    `gorm:"column:created_at"`
	Category    string        `gorm:"column:category"`
	Authors     []*PsqlAuthor `gorm:"many2many:news_authors;joinForeignKey:news_id;joinReferences:authors_id"`
}

func (n *PsqlNewsDoc) TableName() string {
	return "news"
}

// PsqlAuthor is the struct of authors of news docs
type PsqlAuthor struct {
	ID   int    `gorm:"column:id;primarykey"`
	Name string `gorm:"column:name"`
}

func (a *PsqlAuthor) TableName() string {
	return "authors"
}

// PsqlNewsAuthors is the reference table for News and Authors
type PsqlNewsAuthors struct {
	ID         int `gorm:"column:id;primarykey"`
	NewsID     int `gorm:"column:news_id"`
	Authors_ID int `gorm:"column:authors_id"`
}

func (na *PsqlNewsAuthors) TableName() string {
	return "news_authors"
}
