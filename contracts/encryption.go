package contracts

type EncryptorFactory interface {
	Encryptor

	Extend(key string, encryptor Encryptor)

	Driver(key string) Encryptor
}

type Encryptor interface {
	Encode(value string) string
	Decode(encrypted string) (string, error)
}
