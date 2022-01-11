package scheduling

import (
	"github.com/qbhy/goal/contracts"
	"sync"
)

type Schedule struct {
	timezone string
	mutex    sync.Mutex
	app      contracts.Application
}

func (this *Schedule) Call(fn interface{}, args ...interface{}) {
	this.app.Call(fn, args...)
}
