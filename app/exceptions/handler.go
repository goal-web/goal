package exceptions

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/validation"
	"reflect"
	"runtime/debug"
	"strings"
)

type ExceptionHandler struct {
	dontReportExceptions []reflect.Type
}

func NewHandler() contracts.ExceptionHandler {
	return &ExceptionHandler{utils.ConvertToTypes([]contracts.Exception{})}
}

func (handler *ExceptionHandler) Handle(exception contracts.Exception) interface{} {
	logs.WithException(exception).Warn("报错了")
	switch e := exception.(type) {
	case http.Exception: // http 支持在异常处理器返回响应
		return handler.handleHttpException(e)
	case validation.Exception:
		return handler.renderValidationException(e)
	default:
		debug.PrintStack()
	}

	logs.WithException(exception).
		WithField("exception", reflect.TypeOf(exception).String()).
		Error("ExceptionHandler")

	if httpException, isHttpException := exception.(http.Exception); isHttpException {
		logs.WithException(httpException).WithFields(contracts.Fields{}).Debug("http请求报错")
	}

	if handler.ShouldReport(exception) {
		handler.Report(exception)
	}

	return nil
}

func (handler *ExceptionHandler) handleHttpException(exception http.Exception) interface{} {

	switch e := exception.Exception.(type) {
	case validation.Exception:
		return handler.renderValidationException(e)
	default:
		if !strings.Contains(exception.Error(), "404") {
			debug.PrintStack()
		}
		return contracts.Fields{
			"path":  exception.Request.Path(),
			"error": e.Error(),
		}
	}
}

func (handler *ExceptionHandler) renderValidationException(exception validation.Exception) interface{} {
	return contracts.Fields{
		"msg":    exception.Error(),
		"fields": exception.Fields(),
		"errors": exception.GetErrors(),
	}
}

func (handler *ExceptionHandler) Report(exception contracts.Exception) {
}
func (handler *ExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return !utils.IsInstanceIn(exception, handler.dontReportExceptions...)
}
