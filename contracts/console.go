package contracts

type Console interface {
	Call(command string, args ...interface{}) error
}

type Command interface {
	Handle() interface{}
	Signature() string
	Name() string
	Description() string
	Help() string
}
