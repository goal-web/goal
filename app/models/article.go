package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var (
	ArticleClass = class.Make[Article]()
	//ArticleModel = table.NewModel(ArticleClass, "articles")
)

func ArticleQuery() *table.Table[Article] {
	return table.Query[Article]("articles")
}

type Article struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
