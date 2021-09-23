package checkers

import (
	"github.com/qbhy/goal/contracts"
)

// custom 自定义验证
type custom struct {
	checker func(interface{}) error
}

func Custom(checker func(interface{}) error) contracts.Checker {
	return custom{checker}
}

func (this custom) Check(value interface{}) error {
	return this.checker(value)
}
