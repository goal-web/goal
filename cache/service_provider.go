package cache

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
	
}

func (this *ServiceProvider) Register(container contracts.Container) {
	container.ProvideSingleton(func() CacheManager {
		return CacheManager{}
	})
}
