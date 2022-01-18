package tests

import (
	"fmt"
	"github.com/qbhy/goal/database/builder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestSimpleQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users")
	query.Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0, "or").
		OrWhere("amount", ">=", 100).
		WhereIsNull("avatar")
	fmt.Println(query.ToSql())

	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestJoinQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, builder.Or)
	fmt.Println(query.ToSql())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestDistinctQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users").
		Distinct().
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, builder.Or)
	fmt.Println(query.ToSql())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestBetweenQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		WhereFunc(func(b *builder.Builder) {
			// 高瘦
			b.WhereBetween("height", []int{180, 200}).
				WhereBetween("weight", []int{50, 60}).
				WhereIn("id", []int{1, 2, 3, 4, 5})
		}).OrWhereFunc(func(b *builder.Builder) {
		// 矮胖
		b.WhereBetween("height", []int{140, 160}).
			WhereBetween("weight", []int{70, 140}).
			WhereNotBetween("id", []int{1, 5})
	})
	fmt.Println(query.ToSql())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestUnionQueryBuilder(t *testing.T) {
	query := builder.NewQueryBuilder("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, builder.Or).
		Union(
			builder.NewQueryBuilder("peoples"),
		).
		Union(
			builder.NewQueryBuilder("accounts"),
		).
		UnionAll(
			builder.NewQueryBuilder("members"),
		).
		UnionAll(
			builder.NewQueryBuilder("students"),
		)
	fmt.Println(query.ToSql())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestComplexQueryBuilder(t *testing.T) {

	query1 := builder.NewQueryBuilder("users")
	query1.
		FromSub(func() *builder.Builder {
			return builder.NewQueryBuilder("users").Where("amount", ">", 1000)
		}, "rich_users").
		Join("accounts", "users.id", "=", "accounts.user_id").
		WhereFunc(func(b *builder.Builder) {
			b.Where("name", "goal").
				Where("age", "<", "18").
				WhereIn("id", []int{1, 2})
		}).
		OrWhereFunc(func(b *builder.Builder) {
			b.Where("name", "qbhy").
				Where("age", ">", 18).
				WhereNotIn("id", []int{1, 2})
		}).
		OrWhereNotIn("id", []int{6, 7}).
		OrWhereNotNull("id").
		OrderByDesc("age").
		OrderBy("id").
		GroupBy("country")

	_, err := sqlparser.Parse(query1.ToSql())
	assert.Nil(t, err, err)
}
