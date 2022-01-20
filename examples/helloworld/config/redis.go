package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/redis"
	"github.com/goal-web/supports/utils"
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
					"network":  env.GetString("redis.cache.network"),
					"host":     env.GetString("redis.cache.host"),
					"port":     env.GetString("redis.cache.port"),
					"username": env.GetString("redis.cache.username"),
					"password": env.GetString("redis.cache.password"),
					"db":       env.GetInt64("redis.cache.db"),
					"retries":  env.GetInt64("redis.cache.retries"),
				},
			},
		}
	}
}
