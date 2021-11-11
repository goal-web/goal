package routes

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type ServiceProvider struct {
	app contracts.Application
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
	return this.app.Call(func(router contracts.Router, config contracts.Config) error {
		return router.Start(
			utils.StringOr(
				config.GetString("server.address"),
				fmt.Sprintf("%s:%s",
					config.GetString("server.host"),
					utils.StringOr(config.GetString("server.port"), "8000"),
				),
			),
		)
	})[0].(error)
}

func (provider ServiceProvider) Register(container contracts.Application) {
	routerInstance := New(container)
	container.ProvideSingleton(func() contracts.Router {
		return routerInstance
	}, "router")
}
