package contracts

type Fields map[string]interface{}

type FieldsProvider interface {
	Get() Fields
}
