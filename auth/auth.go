package auth

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/utils"
)

type Auth struct {
	config        contracts.Config

	guardDrivers  map[string]contracts.GuardProvider
	guards        map[string]contracts.Guard

	userDrivers   map[string]contracts.UserProviderProvider
	userProviders map[string]contracts.UserProvider
}

func (this *Auth) Once(authorizable contracts.Authorizable) {
	this.Guard("default").Once(authorizable)
}

func (this *Auth) User() contracts.Authorizable {
	return this.Guard("default").User()
}

func (this *Auth) Id() string {
	return this.Guard("default").Id()
}

func (this *Auth) Check() bool {
	return this.Guard("default").Check()
}

func (this *Auth) Guest() bool {
	return this.Guard("default").Guest()
}

func (this *Auth) Validate(credentials contracts.Fields) bool {
	return this.Guard("default").Validate(credentials)
}

func (this *Auth) ExtendUserProvider(key string, provider contracts.UserProviderProvider) {
	this.userDrivers[key] = provider
}

func (this *Auth) ExtendGuard(key string, guard contracts.GuardProvider) {
	this.guardDrivers[key] = guard
}

func (this *Auth) Guard(key string) contracts.Guard {
	if guard, existsGuard := this.guards[key]; existsGuard {
		return guard
	}

	config := this.config.GetFields(fmt.Sprintf("auth.guards.%s", key))
	driver := utils.GetStringField(config, "driver")

	if guardProvider, existsProvider := this.guardDrivers[driver]; existsProvider {
		this.guards[key] = guardProvider(config)
		return this.guards[key]
	}

	panic(Exception{
		Exception: exceptions.New(fmt.Sprintf("不支持的守卫驱动：%s", driver), config),
	})
}

func (this *Auth) UserProvider(key string) contracts.UserProvider {
	if userProvider, existsUserProvider := this.userProviders[key]; existsUserProvider {
		return userProvider
	}

	config := this.config.GetFields(fmt.Sprintf("auth.users.%s", key))
	driver := utils.GetStringField(config, "driver")

	if userDriver, existsProvider := this.userDrivers[driver]; existsProvider {
		this.userProviders[key] = userDriver(config)
		return this.userProviders[key]
	}

	panic(Exception{
		Exception: exceptions.New(fmt.Sprintf("不支持的用户驱动：%s", driver), config),
	})
}
