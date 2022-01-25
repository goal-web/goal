package tests

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/validation"
	"github.com/goal-web/goal/validation/checkers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DemoParam struct {
	Id string
}

func TestValidator(t *testing.T) {
	demoCheckers := contracts.Rules{
		"id": {checkers.StrLen(1, 5)},
	}
	assert.True(t, validation.Make(DemoParam{Id: "55555"}, demoCheckers).IsSuccessful())

	assert.True(t, len(validation.Make(DemoParam{Id: "666666"}, demoCheckers).Errors()) > 0)

	assert.True(t, validation.Make(contracts.Fields{
		"id": "55555",
	}, demoCheckers).IsSuccessful())

	assert.False(t, validation.Make(contracts.Fields{
		"id": "666666",
	}, demoCheckers).IsSuccessful())
}

// 定义一个表单
type DemoForm struct {
	Id       string
	Username string
}

func (d DemoForm) Names() map[string]string {
	return map[string]string{
		"id": "身份证",
	}
}

func (d DemoForm) Validate() contracts.Validator {
	return validation.Make(d, contracts.Rules{
		"id": {checkers.StrLen(1, 5)},
	})
}

//
func TestValidatorForm(t *testing.T) {
	form := DemoForm{Id: "1", Username: "刚好的名字"}
	validator := form.Validate()
	assert.True(t, validator.IsSuccessful())

	fmt.Println(DemoForm{Id: "", Username: "刚好的名字"}.Validate().Errors())
}

// 自定义map校验
func TestValidatorCustomMap(t *testing.T) {
	form := map[string]int{"a": 1}

	result := validation.Make(form, contracts.Rules{
		"a": {checkers.Between(5, 10)},
	}).Names(map[string]string{
		"a": "自定义的A",
	})

	fmt.Println(result.Errors())
	assert.True(t, result.IsFail())
}

// 自定义校验器
func TestValidatorCustomChecker(t *testing.T) {
	form := map[string]int{"a": 1}
	result := validation.Make(form, contracts.Rules{
		"a": {checkers.Between(5, 10)},
	}).Names(map[string]string{
		"a": "自定义的A",
	})
	fmt.Println(result.Errors())
	assert.True(t, result.IsFail())
}

type DemoValidatable struct {
	fields contracts.Fields
}

func (d DemoValidatable) Names() map[string]string {
	return map[string]string{"id": "IDD"}
}

func (d DemoValidatable) Rules() contracts.Rules {
	return map[string][]contracts.Checker{
		"id": {checkers.Custom(func(i interface{}) error {
			if i != nil {
				return nil
			}
			return errors.New("{field}不能为空")
		})},
	}
}

func (d DemoValidatable) Fields() contracts.Fields {
	return d.fields
}

// 自定义校验器
func TestValidatable(t *testing.T) {

	fmt.Println(validation.Valid(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).Validate())

	fmt.Println(validation.Valid(DemoValidatable{fields: map[string]interface{}{
		"id": nil,
	}}).Errors())

	assert.True(t, validation.Valid(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).IsSuccessful())

	assert.True(t, validation.Valid(DemoValidatable{fields: map[string]interface{}{
		"id": nil,
	}}).IsFail())

	validation.Valid(DemoValidatable{fields: map[string]interface{}{
		"id": "不是空的",
	}}).Validate()
}
