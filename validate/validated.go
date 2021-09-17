package validate

import "github.com/qbhy/goal/contracts"

type ValidatedResult struct {
	param  interface{}
	errors contracts.ValidateErrors
}

func (this ValidatedResult) IsFail() bool {
	return len(this.errors) > 0
}

func (this ValidatedResult) IsSuccessful() bool {
	return len(this.errors) == 0
}

func (this ValidatedResult) Errors() contracts.ValidateErrors {
	return this.errors
}

// Assure 如果验证失败就 panic ，保证数据校验结果无异常
func (this ValidatedResult) Assure() {
	if this.IsFail() {
		panic(NewValidatorException(this.param, this.errors))
	}
}
