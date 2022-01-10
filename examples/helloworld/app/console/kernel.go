package console

import (
	"github.com/qbhy/goal/console"
	"github.com/qbhy/goal/console/commands"
)

func GetCommandProviders() []console.CommandProvider {
	return []console.CommandProvider{
		commands.Runner,
	}
}
