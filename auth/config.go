package auth

import "github.com/qbhy/goal/contracts"

type Config struct {
	Defaults struct {
		Guard string
		User  string
	}
	Guards map[string]contracts.GuardProvider
	Users  map[string]contracts.UserProviderProvider
}
