package events

import (
	"github.com/qbhy/goal/contracts"
)

type EventException struct {
	exception contracts.Exception
	event     contracts.Event
}

func (e EventException) Exception() contracts.Exception {
	return e.exception
}

func (e EventException) Error() string {
	return e.exception.Error()
}

func (e EventException) Fields() contracts.Fields {
	return e.exception.Fields()
}
