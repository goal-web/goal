package cache

import (
	"github.com/qbhy/goal/contracts"
	"time"
)

type RedisStore struct {

}

func (r RedisStore) Get(key string) interface{} {
	panic("implement me")
}

func (r RedisStore) Many(keys []string) []interface{} {
	panic("implement me")
}

func (r RedisStore) Put(key string, value interface{}, seconds time.Duration) error {
	panic("implement me")
}

func (r RedisStore) Add(key string, value interface{}, ttl ...time.Duration) error {
	panic("implement me")
}

func (r RedisStore) Pull(key string, defaultValue ...interface{}) interface{} {
	panic("implement me")
}

func (r RedisStore) PutMany(values map[string]interface{}, seconds time.Duration) error {
	panic("implement me")
}

func (r RedisStore) Increment(key string, value ...int64) (int64, error) {
	panic("implement me")
}

func (r RedisStore) Decrement(key string, value ...int64) (int64, error) {
	panic("implement me")
}

func (r RedisStore) Forever(key string, value interface{}) error {
	panic("implement me")
}

func (r RedisStore) Forget(key string) error {
	panic("implement me")
}

func (r RedisStore) Flush() error {
	panic("implement me")
}

func (r RedisStore) GetPrefix() string {
	panic("implement me")
}

func (r RedisStore) Remember(key string, ttl time.Duration, provider contracts.InstanceProvider) interface{} {
	panic("implement me")
}

func (r RedisStore) RememberForever(key string, provider contracts.InstanceProvider) interface{} {
	panic("implement me")
}
