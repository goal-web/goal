package container

import (
	"github.com/goal-web/contracts"
	"reflect"
)

var (
	ComponentType = reflect.TypeOf((*contracts.Component)(nil)).Elem()
)

type Component struct {
}

func (c Component) ShouldInject() {
}
