package database

import (
	"github.com/goal-web/contracts"
)

type Config struct {
	Default     string
	Connections map[string]contracts.Fields
}
