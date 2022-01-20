package builder

import (
	"github.com/goal-web/contracts"
)

type ParamException struct {
	error
	fields contracts.Fields
}

func (p ParamException) Error() string {
	return p.error.Error()
}

func (p ParamException) Fields() contracts.Fields {
	return p.fields
}
