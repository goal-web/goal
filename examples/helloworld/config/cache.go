package config

import (
	"github.com/goal-web/cache"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

func init() {
	configs["cache"] = func(env contracts.Env) interface{} {
		return cache.Config{
			Default: utils.StringOr(env.GetString("cache.default"), "redis"),
			Stores: map[string]contracts.Fields{
				"redis": {
					"driver":     "redis",
					"connection": env.GetString("cache.connection"),
					"prefix":     env.GetString("cache.prefix"),
				},
			},
		}
	}
}
