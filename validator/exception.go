package validator

import "github.com/qbhy/goal/contracts"

type ValidatorException struct {
	errors contracts.ValidateErrors
}

func NewValidatorException(errors contracts.ValidateErrors) ValidatorException {
	return ValidatorException{errors}
}

func (v ValidatorException) Error() (str string) {
	for _, err := range v.errors {
		return err.Error()
	}
	return
}

func (v ValidatorException) Fields() contracts.Fields {
	var fields = make(contracts.Fields, 0)
	for key, err := range v.errors {
		fields[key] = err.Error()
	}
	return fields
}
