package redis

import (
	"fmt"
	goredis "github.com/go-redis/redis/v8"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

type Factory struct {
	config           Config
	exceptionHandler contracts.ExceptionHandler
	connections      map[string]contracts.RedisConnection
}

func (this *Factory) getName(names ...string) string {
	if len(names) > 0 {
		return names[0]
	} else {
		return this.config.Default
	}
}

func (this *Factory) Connection(names ...string) contracts.RedisConnection {
	name := this.getName(names...)

	if connection, existsConnection := this.connections[name]; existsConnection {
		return connection
	}

	config := this.config.Stores[name]

	// todo: 待优化 redis 配置
	this.connections[name] = &Connection{
		exceptionHandler: this.exceptionHandler,
		client: goredis.NewClient(&goredis.Options{
			Network: utils.GetStringField(config, "network", "tcp"),
			Addr: fmt.Sprintf("%s:%s",
				utils.GetStringField(config, "host", "127.0.0.1"),
				utils.GetStringField(config, "port", "6379"),
			),
			Dialer:             nil,
			OnConnect:          nil,
			Username:           utils.GetStringField(config, "username"),
			Password:           utils.GetStringField(config, "password"),
			DB:                 utils.GetIntField(config, "db", 0),
			MaxRetries:         utils.GetIntField(config, "retries", 3),
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        0,
			ReadTimeout:        0,
			WriteTimeout:       0,
			PoolFIFO:           false,
			PoolSize:           0,
			MinIdleConns:       0,
			MaxConnAge:         0,
			PoolTimeout:        0,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			Limiter:            nil,
		}),
	}

	return this.connections[name]
}
