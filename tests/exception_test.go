package tests

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/logs"
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
