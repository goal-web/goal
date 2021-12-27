package session

import (
	"github.com/qbhy/goal/contracts"
)

type RequestBeforeListener struct {
}

func (this *RequestBeforeListener) Handle(event contracts.Event) {
	//TODO implement me
	panic("implement me")
}

type RequestAfterListener struct {
}

func (this *RequestAfterListener) Handle(event contracts.Event) {
	//TODO implement me
	panic("implement me")
}
