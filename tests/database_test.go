package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/database/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getQuery(name string) contracts.QueryBuilder {
	getApp("/Users/qbhy/project/go/goal-web/goal/examples/helloworld")
	return table.WithConnection(name, "sqlite")
}

func TestTableCreate(t *testing.T) {
	user := getQuery("users").Create(contracts.Fields{
		"name": "qbhy",
	}).(contracts.Fields)

	fmt.Println(user)
	assert.True(t, user["id"].(int64) > 0)
}

func TestTableSelect(t *testing.T) {

	users := getQuery("users").Get()

	fmt.Println(users)

	users.Map(func(user contracts.Fields) {
		fmt.Println(user, user["id"])
	})

	assert.True(t, users != nil)
}

func TestTableQuery(t *testing.T) {

	getQuery("users").Delete()

	user := getQuery("users").Create(contracts.Fields{
		"name": "qbhy",
	}).(contracts.Fields)

	fmt.Println(user)
	userId := user["id"].(int64)
	// 判断插入是否成功
	assert.True(t, userId > 0)

	// 获取数据总量
	assert.True(t, getQuery("users").Count() == 1)

	// 修改数据
	num := getQuery("users").Where("name", "qbhy").Update(contracts.Fields{
		"name": "goal",
	})
	assert.True(t, num == 1)
	// 判断修改后的数据
	user = getQuery("users").Where("name", "goal").First().(contracts.Fields)

	assert.True(t, user["id"] == userId)
	assert.True(t, user["name"] == "goal")
	assert.True(t, getQuery("users").Find(userId).(contracts.Fields)["id"] == userId)
	assert.True(t, getQuery("users").Where("id", userId).Delete() == 1)
	assert.Nil(t, getQuery("users").Find(userId))
}
