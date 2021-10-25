package validation

import "github.com/qbhy/goal/contracts"

type ValidatorException struct {
	param  contracts.Fields
	errors contracts.ValidatedErrors
}

func (this ValidatorException) GetParam() contracts.Fields {
	return this.param
}

func NewValidatorException(param contracts.Fields, errors contracts.ValidatedErrors) ValidatorException {
	return ValidatorException{param, errors}
}

func (this ValidatorException) Error() (str string) {
	for _, err := range this.errors {
		if len(err[0]) > 0 {
			return err[0]
		}
	}
	return
}

func (this ValidatorException) Fields() contracts.Fields {
	var fields = make(contracts.Fields, 0)
	for key, err := range this.errors {
		fields[key] = err
	}
	return fields
}
