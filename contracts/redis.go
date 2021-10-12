package contracts

type RedisFactory interface {
	Connection(name ...string) RedisConnection
}

type RedisSubscribeFunc func(message, channel string)

type RedisConnection interface {
	Subscribe(channels []string, closure RedisSubscribeFunc)
	PSubscribe(channels []string, closure RedisSubscribeFunc)
	Command(method string, args ...interface{}) (interface{}, error)
}
