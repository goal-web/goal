package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/views"
)

func init() {
	configs["views"] = func(env contracts.Env) any {
		return views.Config{
			Path: env.StringOptional("views.path", "views"),
		}
	}
}
