package contracts

type Config interface {
	Getter
	FieldsProvider

	Load(provider FieldsProvider)
	Merge(key string, config Config)
	Get(key string, defaultValue ...interface{}) interface{}
	Set(key string, value interface{})
	Unset(key string)
	GetConfig(key string) Config
}

type Env interface {
	Getter
	OptionalGetter

	FieldsProvider

	Load() Fields
}
