package contracts

import "time"

type CacheFactory interface {
	Store(name ...string) CacheStore
}

type CacheRepository interface {
	Pull(key string, defaultValue ...interface{}) interface{}

	Put(key string, value interface{}, ttl ...time.Duration) error

	// Add Store an item in the cache if the key does not exist.
	Add(key string, value interface{}, ttl ...time.Duration) error

	Increment(key string, value ...int64) (int64, error)

	Decrement(key string, value ...int64) (int64, error)

	Forever(key string, value interface{}) error

	Remember(key string, ttl time.Duration, provider InstanceProvider) interface{}

	RememberForever(key string, provider InstanceProvider) interface{}

	Forget(key string) error

	GetStore() CacheStore
}

type CacheStore interface {
	Get(key string) interface{}

	Many(keys []string) []interface{}

	Put(key string, value interface{}, seconds time.Duration) error

	// Add Store an item in the cache if the key does not exist.
	Add(key string, value interface{}, ttl ...time.Duration) error

	Pull(key string, defaultValue ...interface{}) interface{}

	PutMany(values map[string]interface{}, seconds time.Duration) error

	Increment(key string, value ...int64) (int64, error)

	Decrement(key string, value ...int64) (int64, error)

	Forever(key string, value interface{}) error

	Forget(key string) error

	Flush() error

	GetPrefix() string

	Remember(key string, ttl time.Duration, provider InstanceProvider) interface{}

	RememberForever(key string, provider InstanceProvider) interface{}
}
