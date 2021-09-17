package validator

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"reflect"
	"strings"
)

var (
	ValidateTypeError = errors.New("参数类型错误")
)

// Validate 验证参数是否复合规则
func Validate(param interface{}, checkers contracts.Checkers) Validated {
	var (
		validateErrors = make(contracts.ValidateErrors, 0)
		checkField     = func(key string, value interface{}) {
			if fieldCheckers, ok := checkers[key]; ok {
				for _, checker := range fieldCheckers {
					if err := checker.Check(value); err != nil {
						validateErrors[key] = err
						return
					}
				}
			}
		}
	)

	switch mapValue := param.(type) {
	case map[string]interface{}:
		for key, value := range mapValue {
			checkField(key, value)
		}
	case contracts.Fields:
		for key, value := range mapValue {
			checkField(key, value)
		}
	default:
		utils.EachStructField(param, func(field reflect.StructField, value reflect.Value) {
			checkField(strings.ToLower(field.Name), value.Interface())
		})
	}

	return Validated{validateErrors}
}
