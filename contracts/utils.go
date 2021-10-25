package contracts

type Fields map[string]interface{}

type FieldsProvider interface {
	GetFields() Fields
}
