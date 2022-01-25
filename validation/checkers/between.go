package checkers

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/goal/validation"
)

func Between(min, max float64) contracts.Checker {
	return between{
		Message: "",
		Min:     min,
		Max:     max,
	}
}

// between 数字范围
type between struct {
	Message string
	Min     float64
	Max     float64
}

func (this between) SetMessage(message string) contracts.Checker {
	this.Message = message
	return this
}

func (this between) Check(value interface{}) error {
	var num float64
	switch tmpValue := value.(type) {
	case int:
		num = float64(tmpValue)
	case int16:
		num = float64(tmpValue)
	case int8:
		num = float64(tmpValue)
	case int32:
		num = float64(tmpValue)
	case float64:
		num = tmpValue
	case float32:
		num = float64(tmpValue)
	default:
		return validation.ValidateTypeError
	}

	if float64(num) > this.Max || num < this.Min {
		return errors.New(
			utils.StringOr(this.Message, fmt.Sprintf("{field}必须在 %.2f 到 %.2f 之间", this.Min, this.Max)),
		)
	}
	return nil

}
