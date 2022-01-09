package contracts

type Console interface {
	Call(command string, arguments ConsoleArguments) interface{}
	Run(input ConsoleInput) interface{}
}

type Command interface {
	Handle(arguments ConsoleArguments) interface{}
	GetSignature() string
	GetName() string
	GetDescription() string
	GetHelp() string
}

type ConsoleInput interface {
	GetSignature() string
	GetArguments() ConsoleArguments
}

type ConsoleArguments interface {
	FieldsProvider
	Getter
	OptionalGetter
	StringArrayOption(key string, defaultValue []string) []string
	IntArrayOption(key string, defaultValue []int) []int
	Int64ArrayOption(key string, defaultValue []int64) []int64
	FloatArrayOption(key string, defaultValue []float32) []float32
	Float64ArrayOption(key string, defaultValue []float64) []float64
}
