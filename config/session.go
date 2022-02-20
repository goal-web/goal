package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/session"
	"time"
)

func init() {
	configs["session"] = func(env contracts.Env) interface{} {
		return session.Config{
			Driver:     "redis", // 目前支持 cookie、redis
			Encrypt:    true,
			Domain:     env.GetString("session.domain"),
			Lifetime:   time.Duration(env.GetInt("session.lifetime")),
			Connection: "default",         // database、redis 用到
			Key:        "goal_session:%s", // redis 驱动所用到的 key
			Table:      "sessions",        // database 用到
			Name:       env.StringOption("session.name", "goal"),
		}
	}
}
