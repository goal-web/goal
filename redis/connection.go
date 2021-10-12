package redis

import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
)

type Connection struct {
	exceptionHandler contracts.ExceptionHandler
	client           *goredis.Client
}

func (this *Connection) Subscribe(channels []string, closure contracts.RedisSubscribeFunc) {
	go func() {
		pubSub := this.client.Subscribe(context.Background(), channels...)
		defer func(pubSub *goredis.PubSub) {
			err := pubSub.Close()
			if err != nil {
				// 处理异常
				this.exceptionHandler.Handle(exceptions.ResolveException(err))
			}
		}(pubSub)

		pubSubChannel := pubSub.Channel()

		for msg := range pubSubChannel {
			closure(msg.Payload, msg.Channel)
		}
	}()
}

func (this *Connection) PSubscribe(channels []string, closure contracts.RedisSubscribeFunc) {
	go func() {
		pubSub := this.client.PSubscribe(context.Background(), channels...)
		defer func(pubSub *goredis.PubSub) {
			err := pubSub.Close()
			if err != nil {
				// 处理异常
				this.exceptionHandler.Handle(exceptions.ResolveException(err))
			}
		}(pubSub)

		pubSubChannel := pubSub.Channel()

		for msg := range pubSubChannel {
			closure(msg.Payload, msg.Channel)
		}
	}()
}

func (this *Connection) Command(method string, args ...interface{}) (interface{}, error) {
	return this.client.Do(context.Background(), append([]interface{}{method}, args...)...).Result()
}

func (this *Connection) Client() *goredis.Client {
	return this.client
}
