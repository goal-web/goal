package validation

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/utils"
)

func Make(data interface{}, rules contracts.Rules) contracts.Validator {
	fields, err := utils.ConvertToFields(data)
	if err != nil {
		panic(exceptions.New(err.Error(), contracts.Fields{
			"data":  data,
			"rules": rules,
		}))
	}
	validator := &Validator{
		fields:         fields,
		rules:          rules,
		fieldsNamesMap: make(map[string]string),
		isValidated:    false,
		errors:         make(contracts.ValidatedErrors),
	}

	if alias, hasFieldsAlias := data.(contracts.FieldsAlias); hasFieldsAlias {
		validator.Names(alias.Names())
	}

	return validator
}

func Valid(form contracts.ValidatableForm) contracts.Validator {
	validator := &Validator{
		fields:         form.Fields(),
		rules:          form.Rules(),
		fieldsNamesMap: make(map[string]string),
		isValidated:    false,
		errors:         make(contracts.ValidatedErrors),
	}

	if alias, hasFieldsAlias := form.(contracts.FieldsAlias); hasFieldsAlias {
		validator.Names(alias.Names())
	}

	return validator
}
