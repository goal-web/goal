package exceptions

import (
	"github.com/pkg/errors"
	"github.com/qbhy/goal/contracts"
)

// ResolveException 包装 recover 的返回值
func ResolveException(v interface{}) contracts.Exception {
	switch e := v.(type) {
	case contracts.Exception:
		return e
	case error:
		return WithError(e, contracts.Fields{})
	case string:
		return WithError(errors.New(e), contracts.Fields{})
	default:
		return New("error", contracts.Fields{"err": v})
	}
}