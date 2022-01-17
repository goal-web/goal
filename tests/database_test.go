package tests

import (
	"fmt"
	"github.com/qbhy/goal/database/builder"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users")
	query.Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0, "or").
		OrWhere("amount", ">=", 100).
		WhereIsNull("avatar")
	fmt.Println(query.ToSql())

	query1 := builder.NewQueryBuilder("users")
	query1.
		OrWhereFunc(func(b *builder.Builder) {
			b.Where("name", "goal").Where("age", "<", "18").WhereIn("id", []int{1, 2})
		}).
		OrWhereFunc(func(b *builder.Builder) {
			b.Where("name", "qbhy").Where("age", ">", 18).WhereNotIn("id", []int{1, 2})
		}).
		OrWhereNotIn("id", []int{6, 7}).
		OrWhereNotNull("id").
		OrderByDesc("age").
		OrderBy("id")
	fmt.Println(query1.ToSql())
}
