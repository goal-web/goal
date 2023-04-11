package providers

import (
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
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
	"github.com/goal-web/http/sse"
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/goal-web/session"
	"github.com/goal-web/websocket"
	"github.com/golang-module/carbon/v2"
)

type Console struct {
	path string
}

func NewConsole(path string) contracts.ServiceProvider {
	return &App{path}
}

func (app Console) Register(instance contracts.Application) {
	// 设置异常处理器
	instance.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	instance.RegisterServices(
		config.NewService(config.NewDotEnv(config.File("")), config2.GetConfigProviders()),
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystem.ServiceProvider{},
		&serialization.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.NewService(),
		bloomfilter.NewService(),
		auth.NewService(),
		&ratelimiter.ServiceProvider{},
		console.NewService(),
		database.NewService(),
		queue.NewService(false),
		&email.ServiceProvider{},
		&session.ServiceProvider{},
		sse.ServiceProvider{},
		websocket.ServiceProvider{},
		Micro(),
		//&signal.ServiceProvider{},
	)

	instance.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		dispatcher.Register("QUERY_EXECUTED", listeners.DebugQuery{})
	})
}

func (app Console) Start() error {
	return nil
}

func (app Console) Stop() {
}
