package providers

import (
	"github.com/goal-web/auth/gate"
	"github.com/goal-web/contracts"
)

func Gate() contracts.ServiceProvider {
	return &gate.ServiceProvider{}
}
