package exceptions

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
	"reflect"
)

var handler contracts.ExceptionHandler

func init() {
	// 可以手动调用这个方法覆盖异常处理器
	SetExceptionHandler(DefaultExceptionHandler{})
}

func SetExceptionHandler(h contracts.ExceptionHandler) {
	handler = h
}

// ResolveException 包装 recover 的返回值
func ResolveException(v interface{}) Exception {
	switch e := v.(type) {
	case error:
		return WithError(e, contracts.Fields{})
	case contracts.Exception:
		if err, ok := e.(Exception); ok {
			return err
		}
		return Exception{
			err:    e.Error(),
			fields: e.Fields(),
		}
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
	defaultHandler := DefaultExceptionHandler{
		dontReportExceptions: make([]reflect.Type, 0),
	}

	for _, exception := range dontReportExceptions {
		defaultHandler.dontReportExceptions = append(defaultHandler.dontReportExceptions, reflect.TypeOf(exception))
	}

	return defaultHandler
}

func (h DefaultExceptionHandler) Handle(exception contracts.Exception) {
	logs.WithException(exception).Error("DefaultExceptionHandler")
}

func (h DefaultExceptionHandler) Report(exception contracts.Exception) {
}

func (h DefaultExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	for _, t := range h.dontReportExceptions {
		if utils.IsSameStruct(t, exception) {
			return false
		}
	}
	return true
}
