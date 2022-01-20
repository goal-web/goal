package exceptions

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func WithError(err error, fields contracts.Fields) contracts.Exception {
	if e, isException := err.(contracts.Exception); isException {
		return e
	}
	return New(err.Error(), fields)
}

func WithRecover(err interface{}, fields contracts.Fields) contracts.Exception {
	switch e := err.(type) {
	case contracts.Exception:
		return e
	case error:
		return WithError(e, fields)
	case string:
		return New(e, fields)
	case fmt.Stringer:
		return New(e.String(), fields)
	}
	return New(fmt.Sprintf("%v", err), fields)
}

func WithPrevious(err error, fields contracts.Fields, previous error) Exception {
	return Exception{err.Error(), fields, WithError(previous, nil)}
}

func New(err string, fields contracts.Fields) Exception {
	return Exception{err, fields, nil}
}

type Exception struct {
	err      string
	fields   contracts.Fields
	previous contracts.Exception
}

func (e Exception) Error() string {
	return e.err
}

func (e Exception) Fields() contracts.Fields {
	return e.fields
}
