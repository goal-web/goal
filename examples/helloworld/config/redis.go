package config

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/redis"
	"github.com/qbhy/goal/utils"
)

func init() {
	configs["redis"] = func(env contracts.Env) interface{} {
		return redis.Config{
			Default: utils.StringOr(env.GetString("redis.default"), "default"),
			Stores: map[string]contracts.Fields{
				"default": {
					"network":  env.GetString("redis.network"),
					"host":     env.GetString("redis.host"),
					"port":     env.GetString("redis.port"),
					"username": env.GetString("redis.username"),
					"password": env.GetString("redis.password"),
					"db":       env.GetInt64("redis.db"),
					"retries":  env.GetInt64("redis.retries"),
				},
				"cache": {
					"network":  env.GetString("redis.stores.cache.network"),
					"host":     env.GetString("redis.stores.cache.host"),
					"port":     env.GetString("redis.stores.cache.port"),
					"username": env.GetString("redis.stores.caches.username"),
					"password": env.GetString("redis.stores.cache.password"),
					"db":       env.GetInt64("redis.stores.cache.db"),
					"retries":  env.GetInt64("redis.stores.cache.retries"),
				},
			},
		}
	}
}
