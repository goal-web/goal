package console

import (
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/console/commands"
	"github.com/goal-web/supports/logs"
)

var Commands = []contracts.CommandProvider{
	commands.Runner,
	config.EncryptionCommand,
}

func Schedule(schedule contracts.Schedule) {
	schedule.Call(func() {
		//fmt.Println("打印 hello")
		logs.Default().Info("打印 hello")
	}).EveryFiveSeconds()
}
