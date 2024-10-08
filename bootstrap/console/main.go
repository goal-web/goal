package main

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/bootstrap/core"
)

func main() {
	app := core.Application()

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
