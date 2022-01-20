package checkers

import (
	"github.com/goal-web/contracts"
)

// custom 自定义验证
type custom struct {
	checker func(interface{}) error
	message string
}

func Custom(checker func(interface{}) error) contracts.Checker {
	return custom{checker, ""}
}

func (this custom) Check(value interface{}) error {
	return this.checker(value)
}

func (this custom) SetMessage(message string) contracts.Checker {
	message = this.message
	return this
}
