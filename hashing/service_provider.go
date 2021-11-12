package hashing

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type ServiceProvider struct {
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
	return nil
}


func (this ServiceProvider) Register(container contracts.Application) {
	container.Singleton( "hash", func(config contracts.Config) contracts.HasherFactory {
		return &Factory{
			config:  config,
			hashers: make(map[string]contracts.Hasher),
			drivers: map[string]contracts.HasherProvider{
				"bcrypt": func(config contracts.Fields) contracts.Hasher {
					return &Bcrypt{
						cost: utils.GetIntField(config, "cost", 14),
						salt: utils.GetStringField(config, "salt"),
					}
				},
				"md5": func(config contracts.Fields) contracts.Hasher {
					return &Md5{
						salt: utils.GetStringField(config, "salt"),
					}
				},
			},
		}
	})

	container.Singleton("hashing", func(factory contracts.HasherFactory) contracts.Hasher {
		return factory
	})
}
