package contracts

type Checker interface {
	Check(value interface{}) error
	SetMessage(message string) Checker
}

type Rules map[string][]Checker

type ValidatedErrors map[string][]string

// FieldsAlias 有别名
type FieldsAlias interface {
	Names() map[string]string
}

// Validator 验证器
type Validator interface {
	Validate() Fields
	IsFail() bool
	IsSuccessful() bool
	Errors() ValidatedErrors
	Names(names map[string]string) Validator
}

// ValidatableForm 可验证的表单
type ValidatableForm interface {
	FieldsProvider
	Rules() Rules
}