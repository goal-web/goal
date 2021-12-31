package controllers

import (
	"github.com/qbhy/goal/contracts"
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

func DatabaseQuery(db contracts.DBFactory, request contracts.HttpRequest) contracts.Fields {
	connection := db.Connection(request.GetString("connection"))
	var user struct {
		Id   int    `db:"id"`
		Name string `db:"name"`
	}
	err := connection.Get(&user, "select * from users where name=?", "qbhy")
	return contracts.Fields{
		"users": user,
		"err":   err,
	}
}

func RedisExample(redis contracts.RedisConnection) contracts.Fields {
	str, err := redis.Get("incr")
	return contracts.Fields{
		"value": str,
		"err":   err,
	}
}
