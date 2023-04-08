package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/console/inputs"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/providers"
	"os"
)

func main() {
	app := application.Singleton()
	path, _ := os.Getwd()
	app.Instance("path", path)

	app.RegisterServices(providers.NewQueueWorker(path))

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(&inputs.StringArrayInput{ArgsArray: []string{"run"}})
	})
}
