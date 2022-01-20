package filesystem

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/supports/utils"
	"io/fs"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(container contracts.Application) {
	container.Singleton("filesystem", func(config contracts.Config) contracts.FileSystemFactory {
		factory := &Factory{
			config:  config,
			disks:   make(map[string]contracts.FileSystem),
			drivers: make(map[string]contracts.FileSystemProvider),
		}

		factory.Extend("local", func(localConfig contracts.Fields) contracts.FileSystem {
			return &local{
				name: utils.GetStringField(localConfig, "name"),
				root: utils.GetStringField(localConfig, "root"),
				perm: fs.FileMode(utils.GetIntField(localConfig, "perm")),
			}
		})

		return factory
	})

	container.Singleton("system.default", func(factory contracts.FileSystemFactory) contracts.FileSystem {
		return factory
	})
}
