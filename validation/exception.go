package validation

import (
	"github.com/goal-web/contracts"
)

type ValidatorException struct {
	param  contracts.Fields
	errors contracts.ValidatedErrors
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
	return this.param
}

func (this ValidatorException) GetErrors() contracts.ValidatedErrors {
	return this.errors
}
