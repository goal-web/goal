package exceptions

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/http"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
	"github.com/qbhy/goal/validate"
	"reflect"
)

var (
	handler        contracts.ExceptionHandler
	DefaultHandler DefaultExceptionHandler
)

func SetExceptionHandler(h contracts.ExceptionHandler) {
	handler = h
}

// ResolveException 包装 recover 的返回值
func ResolveException(v interface{}) contracts.Exception {
	switch e := v.(type) {
	case contracts.Exception:
		return e
	case error:
		return WithError(e, contracts.Fields{})
	case string:
		return WithError(errors.New(e), contracts.Fields{})
	default:
		return New("error", contracts.Fields{"err": v})
	}
}

// Handle 处理异常
func Handle(exception contracts.Exception) {
	// todo: 加个协程
	defer func() {
		if err := recover(); err != nil {
			logs.WithException(ResolveException(err)).Fatal("异常处理程序出异常了")
		}
	}()
	handler.Handle(exception)
}

type DefaultExceptionHandler struct {
	dontReportExceptions []reflect.Type
}

func NewDefaultHandler(dontReportExceptions []contracts.Exception) DefaultExceptionHandler {
	return DefaultExceptionHandler{utils.ConvertToTypes(dontReportExceptions)}
}

func (h DefaultExceptionHandler) Handle(exception contracts.Exception) {
	switch rawException := exception.(type) {
	case http.HttpException:
		h.HandleHttpException(rawException)
	default:
		logs.WithException(exception).
			WithField("exception", reflect.TypeOf(exception).String()).
			Error("DefaultExceptionHandler")
	}
}

func (h DefaultExceptionHandler) Report(exception contracts.Exception) {
}

func (h DefaultExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return !utils.IsInstanceIn(exception, h.dontReportExceptions...)
}

func (h DefaultExceptionHandler) HandleHttpException(exception http.HttpException) {
	switch preException := exception.Exception.(type) {
	case validate.ValidatorException:
		_ = exception.Context.JSON(400, contracts.Fields{
			"msg":    preException.Error(),
			"errors": preException.Fields(),
			"param":  preException.GetParam(),
		})
	default:
		_ = exception.Context.JSON(500, contracts.Fields{
			"msg":    preException.Error(),
			"errors": preException.Fields(),
		})

		logs.WithException(exception).
			WithField("exception", reflect.TypeOf(exception).String()).
			Error("DefaultExceptionHandler")
	}
}
