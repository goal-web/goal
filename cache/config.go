package cache

import "github.com/qbhy/goal/contracts"

type Config struct {
	Default string
	Stores  map[string]contracts.Fields
}
