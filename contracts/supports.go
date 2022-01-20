package contracts

type Fields map[string]interface{}

type FieldsProvider interface {
	Fields() Fields
}

type Getter interface {
	GetString(key string) string
	GetInt64(key string) int64
	GetInt(key string) int
	GetFloat64(key string) float64
	GetFloat(key string) float32
	GetBool(key string) bool
	GetFields(key string) Fields
}

type OptionalGetter interface {
	StringOption(key string, defaultValue string) string
	Int64Option(key string, defaultValue int64) int64
	IntOption(key string, defaultValue int) int
	Float64Option(key string, defaultValue float64) float64
	FloatOption(key string, defaultValue float32) float32
	BoolOption(key string, defaultValue bool) bool
	FieldsOption(key string, defaultValue Fields) Fields
}
