package tests

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DemoParam struct {
	Id string
}

type StrLen struct {
	Min int
	Max int
}

func (s StrLen) Check(value interface{}) error {
	switch str := value.(type) {
	case string:
		size := len(str)
		if size > s.Max || size < s.Min {
			return errors.New(fmt.Sprintf("字符串长度必须在 %d 到 %d 之间", s.Min, s.Max))
		}
		return nil
	default:
		panic(validator.ValidateTypeError)
	}
}

func TestValidator(t *testing.T) {
	checkers := contracts.Checkers{
		"id": {StrLen{1, 5}},
	}
	assert.True(t, validator.Validate(DemoParam{Id: "55555"}, checkers).IsSuccessful())
	assert.False(t, validator.Validate(DemoParam{Id: "666666"}, checkers).IsSuccessful())

}
