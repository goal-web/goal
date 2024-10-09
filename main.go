package main

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/controllers"
	"github.com/goal-web/goal/bootstrap/core"
	"github.com/goal-web/goal/routes"
	"github.com/goal-web/http"
	"github.com/goal-web/http/sse"
	"github.com/goal-web/http/websocket"
	"github.com/goal-web/session"
	"github.com/goal-web/supports/signal"
	"github.com/goal-web/views"
	"syscall"
)

func main() {
	app := core.Application(core.App{
		QueueWorker:      true,
		SchedulingWorker: true,
	})

	app.RegisterServices(
		views.NewService(),
		http.NewService(
			routes.Api,
			routes.WebSocket,
			routes.Sse,
			controllers.Register,
		),
		session.NewService(),
		sse.NewService(),
		websocket.NewService(),
		signal.NewService(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
	)

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
