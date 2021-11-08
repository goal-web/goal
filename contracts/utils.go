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
