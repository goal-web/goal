package tests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/logs"
	"testing"
)

type DemoExceptionHandler struct {
}

func (d DemoExceptionHandler) Handle(exception contracts.Exception) {
	logs.WithException(exception).Warn("DemoExceptionHandler")
}

func TestExceptionHandler(t *testing.T) {
	exceptions.SetExceptionHandler(DemoExceptionHandler{})

	exceptions.Handle(exceptions.New("报错了", contracts.Fields{
		"id": 1,
	}))
}
