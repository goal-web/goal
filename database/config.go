package database

import "github.com/qbhy/goal/contracts"

type Config struct {
	Default     string
	Connections map[string]contracts.Fields
}
