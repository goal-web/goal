package config

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/cache"
	"github.com/qbhy/goal/supports/utils"
)

func init() {
	configs["cache"] = func(env contracts.Env) interface{} {
		return cache.Config{
			Default: utils.StringOr(env.GetString("cache.default"), "redis"),
			Stores: map[string]ontracts.Fields{
				"redis": {
					"driver":     "redis",
					"connection": env.GetString("cache.connection"),
					"prefix":     env.GetString("cache.prefix"),
				},
			},
		}
	}
}
