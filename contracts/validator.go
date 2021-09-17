package contracts

type Checker interface {
	Check(value interface{}) error
}

type Checkers map[string][]Checker

type ValidateErrors map[string]error

type Validated interface {
	IsFail() bool
	IsSuccessful() bool
	Errors() ValidateErrors
}
