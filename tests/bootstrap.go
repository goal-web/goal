package tests

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/bootstrap/core"
)

func initApp() contracts.Application {
	app := core.Application(core.App{
		QueueWorker:      false,
		SchedulingWorker: false,
	})

	return app
}
