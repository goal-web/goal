package redis

import (
	"fmt"
	goredis "github.com/go-redis/redis/v8"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type Factory struct {
	config           contracts.Config
	exceptionHandler contracts.ExceptionHandler
	connections      map[string]contracts.RedisConnection
}

func (this *Factory) getName(names ...string) string {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = this.config.GetString("cache.default")
	}

	return utils.StringOr(name, "default")
}

func (this Factory) getConfig(name string) contracts.Fields {
	return this.config.GetFields(
		utils.IfString(name == "default", "redis", fmt.Sprintf("redis.stores.%s", name)),
	)
}


func (this *Factory) Connection(names ...string) contracts.RedisConnection {
	name := this.getName(names...)

	if connection, existsConnection := this.connections[name]; existsConnection {
		return connection
	}

	config := this.getConfig(name)

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
			DB:                 int(utils.GetInt64Field(config, "db", 0)),
			MaxRetries:         int(utils.GetInt64Field(config, "retries", 3)),
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
