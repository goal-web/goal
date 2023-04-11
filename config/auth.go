package config

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/goal-web/goal/app/models"
	"github.com/goal-web/supports/class"
	"github.com/golang-jwt/jwt"
	"time"
)

var UserClass = class.Make[contracts.Authenticatable](models.User{})

func init() {
	configs["auth"] = func(env contracts.Env) any {
		return auth.Config{
			Defaults: auth.Defaults{
				Guard: env.StringOptional("auth.default", "jwt"),
				User:  env.StringOptional("auth.user", "db"),
			},
			Guards: map[string]contracts.Fields{
				"jwt": {
					"driver":   "jwt",
					"secret":   env.GetString("auth.jwt.secret"),
					"method":   jwt.SigningMethodHS256,
					"lifetime": 60 * 60 * 24 * time.Second,
					"provider": "db",
				},
				"session": {
					"driver":      "session",
					"provider":    "db",
					"session_key": env.StringOptional("auth.session.key", "auth_session"),
				},
			},
			Users: map[string]contracts.Fields{
				"db": {
					"driver": "db",
					"model":  table.Auth(UserClass, "users", "id"),
				},
			},
		}
	}
}
