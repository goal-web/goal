package exceptions

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
	"reflect"
)

type DefaultExceptionHandler struct {
	dontReportExceptions []reflect.Type
}

func NewDefaultHandler(dontReportExceptions []contracts.Exception) DefaultExceptionHandler {
	return DefaultExceptionHandler{utils.ConvertToTypes(dontReportExceptions)}
}

func (handler DefaultExceptionHandler) Handle(exception contracts.Exception) {
	logs.WithException(exception).
		WithField("exception", reflect.TypeOf(exception).String()).
		Error("DefaultExceptionHandler")
}

func (handler DefaultExceptionHandler) Report(exception contracts.Exception) {
}

func (handler DefaultExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return !utils.IsInstanceIn(exception, handler.dontReportExceptions...)
}
