package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/database/table"
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getQuery(name string) contracts.QueryBuilder {
	getApp("/Users/qbhy/project/go/goal-web/goal/example")
	return table.WithConnection(name, "sqlite")
}
func userModel() contracts.QueryBuilder {
	getApp("/Users/qbhy/project/go/goal-web/goal/example")
	return UserModel().SetConnection("sqlite")
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

	err := getQuery("users").Chunk(10, func(collection contracts.Collection, page int) error {
		assert.True(t, collection.Len() == 1)
		fmt.Println(collection.ToJson())
		return nil
	})

	assert.Nil(t, err)

	assert.True(t, user["id"] == userId)
	assert.True(t, user["name"] == "goal")
	assert.True(t, getQuery("users").Find(userId).(contracts.Fields)["id"] == userId)
	assert.True(t, getQuery("users").Where("id", userId).Delete() == 1)
	assert.Nil(t, getQuery("users").Find(userId))
}

// 定义 class
var UserClass = class.Make(new(User1))

// 定义结构体
type User1 struct {
	Id       int64  `json:"id"`
	NickName string `json:"name"`
}

// 定义模型
func UserModel() *table.Table {
	return table.Model(UserClass, "users")
}

func TestModel(t *testing.T) {
	user := userModel().Create(contracts.Fields{
		"name": "qbhy",
	}).(User1)

	fmt.Println("创建后返回模型", user)

	fmt.Println("用table查询：",
		getQuery("users").Get().ToJson()) // query 返回 Collection<contracts.Fields>

	fmt.Println(userModel(). // model 返回 Collection<User1>
					Get().
					Map(func(user User1) {
			fmt.Println("id:", user.Id)
		}).ToJson())

	fmt.Println(userModel().Where("id", ">", 0).Delete())
}
