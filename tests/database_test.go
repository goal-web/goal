package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/database/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getQuery() contracts.QueryBuilder {
	getApp("/Users/qbhy/project/go/goal-web/goal/examples/helloworld")
	return table.WithConnection("users", "sqlite")
}

func TestTableCreate(t *testing.T) {
	user := getQuery().Create(contracts.Fields{
		"name": "qbhy",
	}).(contracts.Fields)

	fmt.Println(user)
	assert.True(t, user["id"].(int64) > 0)
}

func TestTableSelect(t *testing.T) {

	users := getQuery().Get().([]contracts.Fields)

	fmt.Println(users)

	for _, user := range users {
		fmt.Println(user, user["id"])
	}

	assert.True(t, users != nil)
}

func TestTableQuery(t *testing.T) {

	getQuery().Delete()

	user := getQuery().Create(contracts.Fields{
		"name": "qbhy",
	}).(contracts.Fields)

	fmt.Println(user)
	userId := user["id"].(int64)
	// 判断插入是否成功
	assert.True(t, userId > 0)

	// 获取数据总量
	assert.True(t, getQuery().Count() == 1)

	// 修改数据
	num := getQuery().Where("name", "qbhy").Update(contracts.Fields{
		"name": "goal",
	})
	assert.True(t, num == 1)
	// 判断修改后的数据
	user = getQuery().Where("name", "goal").First().(contracts.Fields)

	assert.True(t, user["id"] == userId)
	assert.True(t, getQuery().Find(userId).(contracts.Fields)["id"] == userId)
	assert.True(t, getQuery().Where("id", userId).Delete() == 1)
	assert.Nil(t, getQuery().Find(userId))
}
