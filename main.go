package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/application/signal"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/goal/app/console"
	"github.com/goal-web/goal/app/exceptions"
	"github.com/goal-web/goal/app/providers"
	config2 "github.com/goal-web/goal/config"
	"github.com/goal-web/goal/database/migrations"
	"github.com/goal-web/goal/routes"
	"github.com/goal-web/hashing"
	"github.com/goal-web/http"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"os"
)

func main() {
	app := application.Singleton()
	path, _ := os.Getwd()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(os.Getenv("env"), path, config2.GetConfigProviders()),
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		//filesystem.serviceProvider{},
		&serialization.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.NewService(),
		bloomfilter.NewService(),
		auth.NewService(),
		//&ratelimiter.serviceProvider{},
		console.NewService(),
		database.Service(migrations.Migrations),
		//&queue.serviceProvider{},
		//&email.serviceProvider{},
		&http.ServiceProvider{RouteCollectors: []interface{}{
			// 路由收集器
			routes.Api,
			routes.WebSocket,
			routes.Sse,
		}},
		//&session.serviceProvider{},
		//sse.serviceProvider{},
		//websocket.serviceProvider{},
		providers.App{},
		//providers.Gate(),
		//providers.Micro(),
		&signal.ServiceProvider{},
	)

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
