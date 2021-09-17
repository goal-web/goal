package checkers

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/validator"
)

// StringLength 字符串长度校验
type StringLength struct {
	Min int
	Max int
}

func (this StringLength) Check(value interface{}) error {
	switch str := value.(type) {
	case string:
		size := len(str)
		if size > this.Max || size < this.Min {
			return errors.New(fmt.Sprintf("字符串长度必须在 %d 到 %d 之间", this.Min, this.Max))
		}
		return nil
	default:
		return validator.ValidateTypeError
	}
}
