package config

import (
	"github.com/qbhy/goal/auth"
	"github.com/goal-web/contracts"
)

func init() {
	configs["auth"] = func(env contracts.Env) interface{} {
		return auth.Config{
			Defaults: struct {
				Guard string
				User  string
			}{
				Guard: env.GetString("auth.default"),
				User:  env.GetString("auth.user"),
			},
			Guards: nil,
			Users:  nil,
		}
	}
}
