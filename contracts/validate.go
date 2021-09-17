package contracts

type Checker interface {
	Check(fieldName string, value interface{}) error
	Message(fieldName string) string
}

type Checkers map[string][]Checker

type ValidateErrors map[string]error

type ValidatedResult interface {
	SafeValidate
	IsFail() bool
	IsSuccessful() bool
	Errors() ValidateErrors
}

type FieldsAlias interface {
	GetFieldAlias(key string) string
}

type Validator interface {
	SafeValidate
	Validate() ValidatedResult
}

type SafeValidate interface {
	Assure()
}

