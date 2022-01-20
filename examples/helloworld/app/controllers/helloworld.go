package controllers

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
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
	err := connection.Get(&user, fmt.Sprintf("select * from users where name='%s'", "qbhy"))
	return contracts.Fields{
		"user": user,
		"err":  err,
	}
}

func RedisExample(redis contracts.RedisConnection) contracts.Fields {
	str, err := redis.Get("incr")
	return contracts.Fields{
		"value": str,
		"err":   err,
	}
}
