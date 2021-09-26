package contracts

type Checker interface {
	Check(value interface{}) error
}

type Checkers map[string][]Checker

type ValidateErrors map[string][]string

type ValidatedResult interface {
	SafeValidate
	IsFail() bool
	IsSuccessful() bool
	Errors() ValidateErrors
}

// FieldsAlias 有别名
type FieldsAlias interface {
	FieldsNameMap() map[string]string
}

// Validator 验证器
type Validator interface {
	Validate() ValidatedResult
}

// SafeValidate 验证不通过即 panic
type SafeValidate interface {
	Assure()
}

// ValidatableForm 可验证的表单
type ValidatableForm interface {
	Checkers() Checkers
	ValidData() Fields
}

// ValidatableAliasForm 可验证并且设置了字段名映射的表单
type ValidatableAliasForm interface {
	FieldsAlias
	ValidatableForm
}
