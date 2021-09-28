package config

import (
	"github.com/qbhy/goal/contracts"
)

type FieldsProvider struct {
	Fields contracts.Fields
}

func (provider FieldsProvider) Get() contracts.Fields {
	return provider.Fields
}

