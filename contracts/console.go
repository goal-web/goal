package contracts

type Console interface {
	Call(command string, arguments CommandArguments) interface{}
	Run(input ConsoleInput) interface{}
}

type Command interface {
	Handle(arguments CommandArguments) interface{}
	GetSignature() string
	GetName() string
	GetDescription() string
	GetHelp() string
}

type ConsoleInput interface {
	GetCommand() string
	GetArguments() CommandArguments
}

type CommandArguments interface {
	FieldsProvider
	Getter
	OptionalGetter
	GetArg(index int) string
	GetArgs() []string
	Exists(key string) bool
	StringArrayOption(key string, defaultValue []string) []string
	IntArrayOption(key string, defaultValue []int) []int
	Int64ArrayOption(key string, defaultValue []int64) []int64
	FloatArrayOption(key string, defaultValue []float32) []float32
	Float64ArrayOption(key string, defaultValue []float64) []float64
}
