package controllers

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/database/drivers"
	"github.com/qbhy/goal/utils"
	"strconv"
)

func HelloWorld() string {
	return "hello, goal."
}

func Counter(session contracts.Session) string {
	count := utils.ConvertToInt(session.Get("count", "0"), 0)
	count++
	session.Put("count", strconv.Itoa(count))
	return "hello, goal." + strconv.Itoa(count)
}

func DatabaseQuery(db contracts.DBConnection) contracts.Fields {
	mysql := db.(*drivers.Mysql)
	var user struct {
		Id   int    `db:"id"`
		Name string `db:"name"`
	}
	err := mysql.Get(&user, "select * from users where name=?", "qbhy")
	return contracts.Fields{
		"users": user,
		"err":   err,
	}
}
