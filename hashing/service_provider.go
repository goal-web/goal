package hashing

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Register(container contracts.Container) {
	container.ProvideSingleton(func(config contracts.Config) contracts.HasherFactory {
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
	}, "hash")

	container.ProvideSingleton(func(factory contracts.HasherFactory) contracts.Hasher {
		return factory
	})
}
