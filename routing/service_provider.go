package routing

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type ServiceProvider struct {
	app contracts.Application

	RouteCollectors []interface{}
}

func (this *ServiceProvider) OnStop() {

}

func (this *ServiceProvider) OnStart() error {
	for _, collector := range this.RouteCollectors {
		this.app.Call(collector)
	}
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

func (this *ServiceProvider) Register(app contracts.Application) {
	this.app = app

	app.ProvideSingleton(func() contracts.Router {
		routerInstance := New(this.app)
		return routerInstance
	}, "router")
}
