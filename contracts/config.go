package contracts

type Config interface {
	Load(provider FieldsProvider)
	Merge(key string, config Config)
	Get(key string, defaultValue ...interface{}) interface{}
	Set(key string, value interface{})
	GetConfig(key string) Config
	GetFields(key string) Fields
	GetString(key string) string
	GetInt(key string) int64
	GetFloat(key string) float64
	GetBool(key string) bool
}
