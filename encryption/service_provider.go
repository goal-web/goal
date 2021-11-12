package encryption

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
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
