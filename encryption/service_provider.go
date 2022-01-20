package encryption

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(container contracts.Application) {
	container.Singleton("encryption", func(config contracts.Config) contracts.EncryptorFactory {
		factory := &Factory{encryptors: make(map[string]contracts.Encryptor)}

		factory.Extend("default", AES(config.GetString("app.key")))

		return factory
	})

	container.Singleton("encryption.default", func(factory contracts.EncryptorFactory) contracts.Encryptor {
		return factory.Driver("default")
	})
}
