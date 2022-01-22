package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"reflect"
	"testing"
)

type DemoExceptionHandler struct {
}

func (d DemoExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return false
}

func (d DemoExceptionHandler) Report(exception contracts.Exception) {
}

func (d DemoExceptionHandler) Handle(exception contracts.Exception) {
	logs.WithException(exception).Warn("DemoExceptionHandler")
}

func TestExceptionHandler(t *testing.T) {
	handler := DemoExceptionHandler{}

	handler.Handle(exceptions.New("报错了", contracts.Fields{
		"id": 1,
	}))

	fmt.Println(reflect.TypeOf(DemoExceptionHandler{}).PkgPath())
}
