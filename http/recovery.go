package http

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/pipeline"
	"github.com/goal-web/supports/exceptions"
)

func (this *Router) recovery(request *Request, next pipeline.Pipe) (result interface{}) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			if res := this.errHandler(panicValue, request); res != nil { // 异常处理器返回的响应优先
				HandleResponse(res, request)
			} else {
				HandleResponse(panicValue, request) // 如果异常处理器没有定义响应，则直接响应 panic 的值
			}
			result = nil
		}
	}()

	// 触发钩子
	this.events.Dispatch(&RequestBefore{request})
	return next(request)
}

func (this *Router) errHandler(err interface{}, request contracts.HttpRequest) (result interface{}) {
	var httpException Exception
	switch rawErr := err.(type) {
	case Exception:
		httpException = rawErr
	default:
		httpException = Exception{
			error:   exceptions.ResolveException(err),
			Request: request,
		}
	}

	// 调用容器内的异常处理器
	return this.app.StaticCall(exceptionHandler, httpException)[0]
}
