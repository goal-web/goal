package checkers

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/validate"
)

// Between 数字范围
type Between struct {
	Min float64
	Max float64
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
		return validate.ValidateTypeError
	}

	if float64(num) > this.Max || num < this.Min {
		return errors.New(fmt.Sprintf("{field}必须在 %.2f 到 %.2f 之间", this.Min, this.Max))
	}
	return nil

}
