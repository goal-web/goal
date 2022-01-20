package validation

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"strings"
)

var (
	ValidateTypeError = errors.New("参数类型错误")
)

type Validator struct {
	fields         contracts.Fields
	rules          contracts.Rules
	fieldsNamesMap map[string]string
	isValidated    bool
	errors         contracts.ValidatedErrors
}

func (this *Validator) IsFail() bool {
	return !this.IsSuccessful()
}

func (this *Validator) IsSuccessful() (result bool) {
	if this.isValidated {
		return len(this.errors) == 0
	}

	defer func() { // 抛异常也就是失败了
		if err := recover(); err != nil {
			result = false
		}
	}()

	this.Validate()

	return len(this.errors) == 0
}

func (this *Validator) Errors() (results contracts.ValidatedErrors) {
	if this.isValidated {
		return this.errors
	}

	defer func() { // 抛异常也就是失败了
		if err := recover(); err != nil {
			if exception, isValidateException := err.(ValidatorException); isValidateException {
				results = exception.errors
			}
		}
	}()

	this.Validate()

	return this.errors
}

// Names 设置字段隐射
func (this *Validator) Names(names map[string]string) contracts.Validator {
	this.fieldsNamesMap = names
	return this
}

func (this *Validator) Validate() contracts.Fields {
	validatedFields := contracts.Fields{}
	for key, value := range this.fields {
		validatedFields[key] = value
		if fieldCheckers, ok := this.rules[key]; ok {
			for _, checker := range fieldCheckers {
				if err := checker.Check(value); err != nil {
					this.errors[key] = append(this.errors[key], strings.ReplaceAll(err.Error(), "{field}", utils.StringOr(this.fieldsNamesMap[key], key)))
				}
			}
		}
	}
	this.isValidated = true

	if len(this.errors) > 0 { // 有错误，抛异常
		panic(NewValidatorException(this.fields, this.errors))
	}

	return validatedFields
}
