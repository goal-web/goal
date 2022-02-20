package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var (
	ArticleClass = class.Make(new(User))
	ArticleModel = table.NewModel(ArticleClass, "articles")
)

func ArticleQuery() *table.Table {
	return table.FromModel(ArticleModel)
}

type Article struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
