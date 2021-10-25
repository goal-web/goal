package checkers

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/qbhy/goal/validation"
)

// Between 数字范围
type Between struct {
	Message string
	Min     float64
	Max     float64
}

func (this Between) SetMessage(message string) contracts.Checker {
	this.Message = message
	return this
}

func (this Between) Check(value interface{}) error {
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
