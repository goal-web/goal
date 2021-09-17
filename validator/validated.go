package validator

import "github.com/qbhy/goal/contracts"

type Validated struct {
	errors contracts.ValidateErrors
}

func (v Validated) IsFail() bool {
	return len(v.errors) > 0
}

func (v Validated) IsSuccessful() bool {
	return len(v.errors) == 0
}

func (v Validated) Errors() contracts.ValidateErrors {
	return v.errors
}

// Assure 如果验证失败就 panic ，保证数据校验结果无异常
func (v Validated) Assure() {
	if v.IsFail() {
		panic(NewValidatorException(v.errors))
	}
}
