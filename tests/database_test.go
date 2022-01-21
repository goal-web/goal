package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/database/table"
	"testing"
)

func TestTableCreateSql(t *testing.T) {
	app := getApp("/Users/qbhy/project/go/goal-web/goal/examples/helloworld")

	fmt.Println(app)

	user := table.Query("users").Create(contracts.Fields{
		"name": "qbhy",
	})

	fmt.Println(user)
}
