package redis

import (
	"github.com/goal-web/contracts"
)

type Config struct {
	Default string
	Stores  map[string]contracts.Fields
}
