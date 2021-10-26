package exceptions

import (
	"github.com/qbhy/goal/contracts"
)

func WithError(err error, fields contracts.Fields) Exception {
	return New(err.Error(), fields)
}

func New(err string, fields contracts.Fields) Exception {
	return Exception{err, fields}
}

type Exception struct {
	err    string
	fields contracts.Fields
}

func (e Exception) Error() string {
	return e.err
}

func (e Exception) GetFields() contracts.Fields {
	return e.fields
}
