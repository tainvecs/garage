package newssvc

import "time"

// PsqlNewsDoc is the struct of news docs stored in PostgreSQL database
type News struct {
	ID          int        `gorm:"column:id;primarykey"`
	UUID        string     `gorm:"column:uuid"`
	Link        string     `gorm:"column:link"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	Category    string     `gorm:"column:category"`
	Authors     []*Author  `gorm:"many2many:news_authors;joinForeignKey:news_id;joinReferences:authors_id"`
}

func (n *News) TableName() string {
	return "news"
}

// Author is the struct of authors of news docs
type Author struct {
	ID   int    `gorm:"column:id;primarykey"`
	Name string `gorm:"column:name"`
}

func (a *Author) TableName() string {
	return "authors"
}

// NewsAuthors is the reference table for News and Authors
type NewsAuthors struct {
	ID         int `gorm:"column:id;primarykey"`
	NewsID     int `gorm:"column:news_id"`
	Authors_ID int `gorm:"column:authors_id"`
}

func (na *NewsAuthors) TableName() string {
	return "news_authors"
}
