package config

import (
	"time"

	"github.com/goal-web/cache"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

func init() {
	configs["cache"] = func(env contracts.Env) any {
		return cache.Config{
			Default: utils.StringOr(env.GetString("cache.default"), "memory"),
			Stores: map[string]contracts.Fields{
				"memory": {
					"driver": "ram",
					"prefix": env.GetString("cache.prefix"),
					"ttl":    utils.IntOr(env.GetInt("cache.ttl"), 24*int(time.Hour)), // 默认缓存生命周期
				},
				"file": {
					"driver": "file",
					"path":   utils.StringOr(env.GetString("cache.file.path"), "./storage/cache"),
					"prefix": env.GetString("cache.prefix"),
				},
				"redis": {
					"driver":     "redis",
					"connection": utils.StringOr(env.GetString("cache.connection"), "default"),
					"prefix":     env.GetString("cache.prefix"),
				},
			},
		}
	}
}
