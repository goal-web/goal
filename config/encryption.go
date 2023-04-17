package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/encryption"
)

func init() {
	configs["encryption"] = func(env contracts.Env) any {
		return encryption.Config{
			Default: "AES",
			Drivers: map[string]contracts.EncryptDriver{},
		}
	}
}
