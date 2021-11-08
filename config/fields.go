package config

import (
	"github.com/qbhy/goal/contracts"
)

type FieldsProvider struct {
	Data contracts.Fields
}

func (provider FieldsProvider) Fields() contracts.Fields {
	return provider.Data
}

