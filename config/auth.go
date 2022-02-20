package config

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
	"github.com/golang-jwt/jwt"
)

func init() {
	configs["auth"] = func(env contracts.Env) interface{} {
		return auth.Config{
			Defaults: auth.Defaults{
				Guard: env.StringOption("auth.default", "jwt"),
				User:  env.StringOption("auth.user", "db"),
			},
			Guards: map[string]contracts.Fields{
				"jwt": {
					"driver":   "jwt",
					"secret":   env.GetString("auth.jwt.secret"),
					"method":   jwt.SigningMethodHS256,
					"lifetime": 60 * 60 * 24, // 单位：秒
					"provider": "db",
				},
				"session": {
					"driver":      "session",
					"provider":    "db",
					"session_key": env.StringOption("auth.session.key", "auth_session"),
				},
			},
			Users: map[string]contracts.Fields{
				"db": {
					"driver": "db",
					"model":  models.UserModel,
				},
			},
		}
	}
}
