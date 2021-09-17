package validate

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"reflect"
)

var (
	ValidateTypeError = errors.New("参数类型错误")
)

// Make 验证参数是否复合规则
func Make(param interface{}, checkers contracts.Checkers) *Validator {
	return &Validator{param, checkers}
}

type Validator struct {
	param    interface{}
	checkers contracts.Checkers
}

// Assure 如果验证失败就 panic ，保证数据校验结果无异常
func (this Validator) Assure() {
	validated := this.Validate()

	if validated.IsFail() {
		panic(NewValidatorException(this.param, validated.Errors()))
	}
}

func (this Validator) Validate() contracts.ValidatedResult {
	var (
		validateErrors = make(contracts.ValidateErrors, 0)
		checkField     = func(fieldName, key string, value interface{}) {
			if fieldCheckers, ok := this.checkers[key]; ok {
				for _, checker := range fieldCheckers {
					if err := checker.Check(fieldName, value); err != nil {
						validateErrors[key] = err
						return
					}
				}
			}
		}
	)

	switch paramValue := this.param.(type) {
	case map[string]interface{}:
		for key, value := range paramValue {
			checkField(key, key, value)
		}
	case contracts.Fields:
		for key, value := range paramValue {
			checkField(key, key, value)
		}
	case contracts.FieldsAlias:
		utils.EachStructField(paramValue, func(field reflect.StructField, value reflect.Value) {
			name := utils.SnakeString(field.Name)
			checkField(
				utils.StringOr(paramValue.GetFieldAlias(name), field.Name),
				name,
				value.Interface(),
			)
		})
	default:
		utils.EachStructField(this.param, func(field reflect.StructField, value reflect.Value) {
			checkField(field.Name, utils.SnakeString(field.Name), value.Interface())
		})
	}

	return ValidatedResult{this.param,validateErrors}
}
