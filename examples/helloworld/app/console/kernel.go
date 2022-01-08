package console

import (
	"github.com/qbhy/goal/console"
	"github.com/qbhy/goal/console/commonds"
)

func GetCommandProviders() []console.CommandProvider {
	return []console.CommandProvider{
		commonds.Runner,
	}
}
