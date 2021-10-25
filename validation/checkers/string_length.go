package checkers

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/qbhy/goal/validation"
)

// StringLength 字符串长度校验
type StringLength struct {
	Min     int
	Max     int
	Message string
}

func StrLen(min, max int) StringLength {
	return StringLength{
		Min:     min,
		Max:     max,
		Message: "",
	}
}

func (this StringLength) SetMessage(message string) contracts.Checker {
	this.Message = message
	return this
}

func (this StringLength) Check(value interface{}) error {
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
