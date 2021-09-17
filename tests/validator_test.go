package tests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/validate"
	"github.com/qbhy/goal/validate/checkers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DemoParam struct {
	Id string
}

func TestValidator(t *testing.T) {
	demoCheckers := contracts.Checkers{
		"id": {checkers.StringLength{1, 5}},
	}
	assert.True(t, validate.Make(DemoParam{Id: "55555"}, demoCheckers).Validate().IsSuccessful())
	assert.False(t, validate.Make(DemoParam{Id: "666666"}, demoCheckers).Validate().IsSuccessful())

	assert.True(t, validate.Make(contracts.Fields{
		"id": "55555",
	}, demoCheckers).Validate().IsSuccessful())

	assert.False(t, validate.Make(contracts.Fields{
		"id": "666666",
	}, demoCheckers).Validate().IsSuccessful())
}

// 定义一个表单
type DemoForm struct {
	Id       string
	Username string
}

func (d DemoForm) GetFieldAlias(key string) string {
	switch key {
	case "username":
		return "用户名"
	}
	return "" // 返回默认的
}

func (d DemoForm) Validate() contracts.ValidatedResult {
	return validate.Make(d, contracts.Checkers{
		"id":       {checkers.StringLength{1, 5}},
		"username": {checkers.StringLength{1, 5}},
	}).Validate()
}

//
func TestValidatorForm(t *testing.T) {
	form := DemoForm{Id: "1", Username: "刚好的名字"}
	result := form.Validate()
	assert.True(t, result.IsSuccessful())
}
