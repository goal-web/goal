package encryption

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/supports/utils"
)

type Factory struct {
	encryptors map[string]contracts.Encryptor
}

func (this *Factory) Encode(value string) string {
	return this.Driver("default").Encode(value)
}

func (this *Factory) Decode(payload string) (string, error) {
	return this.Driver("default").Decode(payload)
}

func (this *Factory) Extend(key string, encryptor contracts.Encryptor) {
	this.encryptors[key] = encryptor
}

func (this *Factory) Driver(key string) contracts.Encryptor {
	return this.encryptors[utils.StringOr(key, "default")]
}
