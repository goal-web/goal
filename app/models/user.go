package models

import (
	"github.com/goal-web/auth/gate"
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var UserClass = class.Make[User]()

func UserQuery() *table.Table[User] {
	return table.Class(UserClass, "users")
}

type Settings struct {
	XxxSwitch bool `json:"xxx_switch"`
}

type User struct {
	Id       string   `json:"id"` // 没有 db tag 默认解析 json
	NickName string   `json:"name"`
	Role     string   `json:"role"`
	Settings Settings `json:"settings"`
}

// Can 实现 gate 需要的方法
func (u User) Can(ability string, arguments ...any) bool {
	return gate.Check(u, ability, arguments...)
}

// GetId 实现 auth 需要的方法
func (u User) GetId() string {
	return u.Id
}
