package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console/inputs"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/goal/app/console"
	"github.com/goal-web/goal/app/exceptions"
	"github.com/goal-web/goal/app/listeners"
	config2 "github.com/goal-web/goal/config"
	"github.com/goal-web/hashing"
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/goal-web/supports/utils"
	"github.com/golang-module/carbon/v2"
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
		config.NewService(utils.StringOr(os.Getenv("env"), "local"), path, config2.GetConfigProviders()),
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystem.ServiceProvider{},
		&serialization.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.NewService(),
		bloomfilter.NewService(),
		//auth.NewService(),
		&ratelimiter.ServiceProvider{},
		console.NewService(),
		database.NewService(),
		&email.ServiceProvider{},
		queue.NewService(true),
		//&signal.ServiceProvider{},
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		dispatcher.Register("QUERY_EXECUTED", listeners.DebugQuery{})
	})

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(&inputs.StringArrayInput{ArgsArray: []string{"run"}})
	})
}
