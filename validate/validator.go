package validate

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

// Make 验证参数是否复合规则
func Make(param interface{}, args ...interface{}) *Validator {
	switch paramValue := param.(type) {
	case contracts.ValidatableAliasForm:
		return &Validator{
			paramValue.GetFields(),
			paramValue.GetCheckers(),
			paramValue.GetFieldsNameMap(),
		}
	case contracts.ValidatableForm:
		return &Validator{
			paramValue.GetFields(),
			paramValue.GetCheckers(),
			make(map[string]string, 0),
		}
	default:
		if len(args) > 0 {
			if checkers, isCheckers := args[0].(contracts.Checkers); isCheckers {
				return &Validator{
					paramValue,
					checkers,
					make(map[string]string, 0),
				}
			}
		}
		panic(errors.New("未设置校验规则"))
	}
}

type Validator struct {
	fields         interface{}
	checkers       contracts.Checkers
	fieldsNamesMap map[string]string
}

// Assure 如果验证失败就 panic ，保证数据校验结果无异常
func (this Validator) Assure() {
	validated := this.Validate()

	if validated.IsFail() {
		panic(NewValidatorException(this.fields, validated.Errors()))
	}
}

// SetFieldsNameMap 设置字段隐射
func (this *Validator) SetFieldsNameMap(names map[string]string) *Validator {
	this.fieldsNamesMap = names
	return this
}

func (this *Validator) Validate() contracts.ValidatedResult {
	var (
		validateErrors    = make(contracts.ValidateErrors, 0)
		checkFieldHandler = func(key string, value interface{}) {
			if fieldCheckers, ok := this.checkers[key]; ok {
				for _, checker := range fieldCheckers {
					if err := checker.Check(value); err != nil {
						validateErrors[key] = append(validateErrors[key], strings.ReplaceAll(err.Error(), "{field}", utils.StringOr(this.fieldsNamesMap[key], key)))
					}
				}
			}
		}
	)

	switch paramValue := this.fields.(type) {
	case map[string]interface{}:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	case map[string]int:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	case map[string]float64:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	case map[string]string:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	case map[string]bool:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	case contracts.Fields:
		for key, value := range paramValue {
			checkFieldHandler(key, value)
		}
	default:
		paramType := reflect.ValueOf(this.fields)

		switch paramType.Kind() {
		case reflect.Struct: // 结构体
			utils.EachStructField(this.fields, func(field reflect.StructField, value reflect.Value) {
				checkFieldHandler(utils.SnakeString(field.Name), value.Interface())
			})
		case reflect.Map: // 自定义的 map
			if paramType.Type().Key().Kind() == reflect.String {
				for _, key := range paramType.MapKeys() {
					name := key.String()
					checkFieldHandler(name, paramType.MapIndex(key).Interface())
				}
			} else {
				panic(errors.New("不支持非string以外的类型作为key的map"))
			}
		default:
			panic(errors.New("不支持验证的类型 " + paramType.String()))
		}
	}

	return ValidatedResult{this.fields, validateErrors}
}
