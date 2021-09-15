package main

import (
	"github.com/labstack/echo/v4"
	appExceptions "qbhy/app/exceptions"
	"qbhy/contracts"
	"qbhy/events"
	"qbhy/exceptions"
	"qbhy/logs"
	"qbhy/http"
)

func main() {
	// 注册异常处理器
	exceptions.SetExceptionHandler(appExceptions.ExceptionHandler{})

	// 注册事件监听器
	events.SetEventListeners(map[contracts.EventName][]contracts.EventListener{})

	router := http.New()

	router.Get("/", func(context echo.Context) interface{} {
		panic("控制器panic")

		return "返回了啥"
	})

	logs.WithError(router.Start(":8000")).Debug("服务器报错了")
}
