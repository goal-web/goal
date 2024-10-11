package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
)

func init() {
	configs["http"] = func(env contracts.Env) any {
		config := http.Config{
			Host:              env.GetString("http.host"),
			Port:              env.GetString("http.port"),
			StaticDirectories: map[string]string{
				//"/": "public",
			},
		}

		if env.GetString("app.env") == "local" {
			config.StaticDirectories["/"] = "public"
		}

		return config
	}
}
