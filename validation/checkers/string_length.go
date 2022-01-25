package checkers

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/goal/validation"
)

// stringLength 字符串长度校验
type stringLength struct {
	Min     int
	Max     int
	Message string
}

func StrLen(min, max int) contracts.Checker {
	return stringLength{
		Min:     min,
		Max:     max,
		Message: "",
	}
}

func (this stringLength) SetMessage(message string) contracts.Checker {
	this.Message = message
	return this
}

func (this stringLength) Check(value interface{}) error {
	switch str := value.(type) {
	case string:
		size := len([]rune(str))
		if size > this.Max || size < this.Min {
			return errors.New(
				utils.StringOr(this.Message, fmt.Sprintf("{field}的长度必须在 %d 到 %d 之间", this.Min, this.Max)),
			)
		}
		return nil
	default:
		return validation.ValidateTypeError
	}
}
