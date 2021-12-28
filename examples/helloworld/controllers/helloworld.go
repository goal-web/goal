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
