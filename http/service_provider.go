package http

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
	"net/http"
)

type ServiceProvider struct {
	app contracts.Application

	RouteCollectors []interface{}
}

func (this *ServiceProvider) Stop() {
	this.app.Call(func(dispatcher contracts.EventDispatcher, router contracts.Router) {
		if err := router.Close(); err != nil {
			logs.WithError(err).Info("router 关闭报错")
		}
		dispatcher.Dispatch(&HttpServeClosed{})
	})
}

func (this *ServiceProvider) Start() error {
	for _, collector := range this.RouteCollectors {
		this.app.Call(collector)
	}

	err := this.app.Call(func(router contracts.Router, config contracts.Config) error {
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

	if err != nil && err != http.ErrServerClosed {
		logs.WithError(err).Error("http 服务无法启动")
		this.app.Stop()
		return err
	}

	return nil
}

func (this *ServiceProvider) Register(app contracts.Application) {
	this.app = app

	app.Singleton("router", func() contracts.Router {
		return New(this.app)
	})
}
