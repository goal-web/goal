package main

import (
	"github.com/goal-web/goal/bootstrap/core"
	"github.com/goal-web/supports/logs"
)

func main() {
	app := core.Application(core.App{
		QueueWorker: true,
	})

	if errors := app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Fatal("goal 异常!")
	} else {
		logs.Default().Info("goal 已关闭")
	}
}
