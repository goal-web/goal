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
func Validate(param interface{}, checkers contracts.Checkers) contracts.Validated {
	var validateErrors = make(contracts.ValidateErrors, 0)

	utils.EachStructField(param, func(field reflect.StructField, value reflect.Value) {
		name := strings.ToLower(field.Name)
		if fieldCheckers, ok := checkers[name]; ok {
			for _, checker := range fieldCheckers {
				if err := checker.Check(value.Interface()); err != nil {
					validateErrors[name] = err
					return
				}
			}
		}
	})

	return Validated{validateErrors}
}
