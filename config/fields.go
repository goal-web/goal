package config

import (
	"github.com/goal-web/contracts"
)

type FieldsProvider struct {
	Data contracts.Fields
}

func (provider FieldsProvider) Fields() contracts.Fields {
	return provider.Data
}
