package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console/inputs"
	"github.com/goal-web/console/scheduling"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/goal/app/console"
	"github.com/goal-web/goal/app/exceptions"
	"github.com/goal-web/goal/app/providers"
	config2 "github.com/goal-web/goal/config"
	"github.com/goal-web/goal/routes"
	"github.com/goal-web/hashing"
	"github.com/goal-web/http"
	"github.com/goal-web/http/sse"
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/goal-web/session"
	"github.com/goal-web/supports/signal"
	"github.com/goal-web/websocket"
	"github.com/golang-module/carbon/v2"
	"syscall"
)

func main() {
	env := config.NewToml(config.File("config.toml"))
	app := application.Singleton(env.GetBool("app.debug"))

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(env, config2.GetConfigProviders()),
		hashing.NewService(),
		encryption.NewService(),
		filesystem.NewService(),
		serialization.NewService(),
		events.NewService(),
		providers.NewEvents(),
		redis.NewService(),
		cache.NewService(),
		bloomfilter.NewService(),
		auth.NewService(),
		ratelimiter.NewService(),
		console.NewService(),
		scheduling.NewService(),
		database.NewService(),
		queue.NewService(true),
		email.NewService(),
		http.NewService(routes.Api, routes.WebSocket, routes.Sse),
		session.NewService(),
		sse.NewService(),
		websocket.NewService(),
		providers.NewMicro(true),
		signal.NewService(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher, console3 contracts.Console) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		console3.Run(inputs.String("run"))
	})
}
