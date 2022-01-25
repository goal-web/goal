package config

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/http"
)

func init() {
	configs["http"] = func(env contracts.Env) interface{} {
		return http.Config{
			Host: env.GetString("http.host"),
			Port: env.GetString("http.port"),
		}
	}
}
