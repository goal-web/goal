package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var (
	ArticleClass = class.Make[Article]()
)

func Articles() *table.Table[Article] {
	return table.Class(ArticleClass, "articles").SetPrimaryKey("id")
}

type Article struct {
	table.Model[Article] `json:"-"`

	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}
