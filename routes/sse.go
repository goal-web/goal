package routes

import (
	"github.com/goal-web/contracts"
	sse2 "github.com/goal-web/goal/app/http/sse"
	"github.com/goal-web/http/sse"
)

func Sse(router contracts.Router) {
	// 自定义 sse 控制器
	router.Get("/sse-demo", sse.New(sse2.DemoController{}))

	// 默认 sse 控制器
	router.Get("/sse", sse.Default())

	router.Get("/send-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Send(uint64(request.GetInt64("fd")), request.GetString("msg"))
	})

	router.Get("/close-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Close(uint64(request.GetInt64("fd")))
	})
}
