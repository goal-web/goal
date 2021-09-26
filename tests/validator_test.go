package tests

import (
	"errors"
	"fmt"
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

func (d DemoForm) FieldsNameMap() map[string]string {
	return map[string]string{
		"id": "身份证",
	}
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

	fmt.Println(DemoForm{Id: "", Username: "刚好的名字"}.Validate().Errors())
}

// 自定义map校验
func TestValidatorCustomMap(t *testing.T) {
	form := map[string]int{"a": 1}
	result := validate.Make(form, contracts.Checkers{
		"a": {checkers.Between{5, 10}},
	}).SetFieldsNameMap(map[string]string{
		"a": "自定义的A",
	}).Validate()
	fmt.Println(result.Errors())
	assert.True(t, result.IsFail())
}

// 自定义校验器
func TestValidatorCustomChecker(t *testing.T) {
	form := map[string]int{"a": 1}
	result := validate.Make(form, contracts.Checkers{
		"a": {checkers.Between{5, 10}},
	}).SetFieldsNameMap(map[string]string{
		"a": "自定义的A",
	}).Validate()
	fmt.Println(result.Errors())
	assert.True(t, result.IsFail())
}

type DemoValidatable struct {
	fields contracts.Fields
}

func (d DemoValidatable) FieldsNameMap() map[string]string {
	return map[string]string{"id": "IDD"}
}

func (d DemoValidatable) Checkers() contracts.Checkers {
	return map[string][]contracts.Checker{
		"id": {checkers.Custom(func(i interface{}) error {
			if i != nil {
				return nil
			}
			return errors.New("{field}不能为空")
		})},
	}
}

func (d DemoValidatable) ValidData() contracts.Fields {
	return d.fields
}

// 自定义校验器
func TestValidatable(t *testing.T) {

	fmt.Println(validate.Make(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).Validate().Errors())

	fmt.Println(validate.Make(DemoValidatable{fields: map[string]interface{}{
		"id": nil,
	}}).Validate().Errors())

	assert.True(t, validate.Make(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).Validate().IsSuccessful())

	assert.True(t, validate.Make(DemoValidatable{fields: map[string]interface{}{
		"id": nil,
	}}).Validate().IsFail())

	validate.Make(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).Validate().Assure()
}
