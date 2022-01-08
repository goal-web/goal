package console

import (
	"github.com/qbhy/goal/console"
	"github.com/qbhy/goal/console/commonds"
)

var commandProviders []console.CommandProvider

func init() {
	commandProviders = []console.CommandProvider{
		commonds.Runner,
	}
}

func GetCommandProviders() []console.CommandProvider {
	return commandProviders
}
