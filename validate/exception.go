package validate

import "github.com/qbhy/goal/contracts"

type ValidatorException struct {
	param  interface{}
	errors contracts.ValidateErrors
}

func (this ValidatorException) GetParam() interface{} {
	return this.param
}

func NewValidatorException(param interface{}, errors contracts.ValidateErrors) ValidatorException {
	return ValidatorException{param, errors}
}

func (this ValidatorException) Error() (str string) {
	for _, err := range this.errors {
		return err.Error()
	}
	return
}

func (this ValidatorException) Fields() contracts.Fields {
	var fields = make(contracts.Fields, 0)
	for key, err := range this.errors {
		fields[key] = err.Error()
	}
	return fields
}
