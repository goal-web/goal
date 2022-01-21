package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestSimpleQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users")
	query.Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0, contracts.Or).
		OrWhere("amount", ">=", 100).
		WhereIsNull("avatar")
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())

	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestJoinQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		JoinSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").
				Where("level", ">", 5)
		}, "vip_users", "vip_users.id", "=", "users.id").
		//WhereIn("gender", "1,2").
		WhereIn("gender", []int{1, 2})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestFromSubQueryBuilder(t *testing.T) {
	query := builder.FromSub(func() contracts.QueryBuilder {
		return builder.NewQuery("users").
			Where("level", ">", 5)
	}, "vip_users").
		//WhereIn("gender", "1,2").
		WhereIn("gender", []int{1, 2})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestDistinctQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Distinct().
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestUpdateSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).UpdateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestSelectSub(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).
		SelectSub(func() contracts.QueryBuilder {
			return builder.NewQuery("accounts").Where("accounts.id", "users.id").WithCount()
		}, "accounts_count").
		Join("accounts", "accounts.user_id", "=", "users.id").
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestWhereNotExists(t *testing.T) {
	sql, bindings := builder.NewQuery("users").
		Where("id", ">", 1).
		WhereNotExists(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Select("id").Where("age", ">", 18)
		}).
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestCount(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).WithCount("id").SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestDeleteSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertIgnoreSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertIgnoreSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertReplaceSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertReplaceSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestCreateSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").CreateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestBetweenQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		WhereFunc(func(b contracts.QueryBuilder) {
			// 高瘦
			b.WhereBetween("height", []int{180, 200}).
				WhereBetween("weight", []int{50, 60}).
				WhereIn("id", []int{1, 2, 3, 4, 5})
		}).OrWhereFunc(func(b contracts.QueryBuilder) {
		// 矮胖
		b.WhereBetween("height", []int{140, 160}).
			WhereBetween("weight", []int{70, 140}).
			WhereNotBetween("id", []int{1, 5})
	})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestUnionQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or).
		UnionByProvider(
			func() contracts.QueryBuilder {
				return builder.NewQuery("peoples").Where("id", 5)
			},
		).
		Union(
			builder.NewQuery("accounts"),
		).
		UnionAll(
			builder.NewQuery("members"),
		).
		UnionAll(
			builder.NewQuery("students"),
		)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestComplexQueryBuilder(t *testing.T) {

	query := builder.NewQuery("users")
	query.
		FromSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Where("amount", ">", 1000)
		}, "rich_users").
		Join("accounts", "users.id", "=", "accounts.user_id").
		WhereFunc(func(b contracts.QueryBuilder) {
			b.Where("name", "goal").
				Where("age", "<", "18").
				WhereIn("id", []int{1, 2})
		}).
		OrWhereFunc(func(b contracts.QueryBuilder) {
			b.Where("name", "qbhy").
				Where("age", ">", 18).
				WhereNotIn("id", []int{1, 2})
		}).
		OrWhereNotIn("id", []int{6, 7}).
		OrWhereNotNull("id").
		OrderByDesc("age").
		OrderBy("id").
		GroupBy("country")

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestGroupByQueryBuilder(t *testing.T) {

	query := builder.
		FromSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Where("amount", ">", 1000)
		}, "rich_users").
		GroupBy("country").
		Having("count(rich_users.id)", "<", 1000).   // 人口少
		OrHaving("sum(rich_users.amount)", "<", 100) // 或者穷

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}
