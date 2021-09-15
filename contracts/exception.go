package contracts

type Exception interface {
	error
	Fields() Fields
}

type ExceptionHandler interface {
	Handle(exception Exception)
}
