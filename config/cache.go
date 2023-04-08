package config

import (
	"github.com/goal-web/cache"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"time"
)

func init() {
	configs["cache"] = func(env contracts.Env) any {
		return cache.Config{
			Default: utils.StringOr(env.GetString("cache.default"), "redis"),
			Stores: map[string]contracts.Fields{
				"memory": {
					"driver": "memory",
					"prefix": env.GetString("cache.prefix"),
					"ttl":    24 * int(time.Hour), // 默认缓存生命周期
				},
				"redis": {
					"driver":     "redis",
					"connection": env.GetString("cache.connection"),
					"prefix":     env.GetString("cache.prefix"),
				},
			},
		}
	}
}
