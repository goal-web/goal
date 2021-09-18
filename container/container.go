package container

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"reflect"
)

var (
	CallerTypeError   = errors.New("参数类型必须是有一个返回值的函数")
	ArgumentsLenError = errors.New("参数长度不一致")
)

type Container struct {
	binds      map[string]contracts.InstanceProvider
	singletons map[string]contracts.InstanceProvider
	instances  map[string]interface{}
	aliases    map[string]string
}

func (this *Container) Provide(provider interface{}) {
	providerType := reflect.TypeOf(provider)
	if providerType.Kind() != reflect.Func || providerType.NumOut() != 1 {
		panic(CallerTypeError)
	}
	resultType := providerType.Out(0)
	this.Bind(utils.GetTypeKey(resultType), func() interface{} {
		return reflect.ValueOf(provider).Call(nil)[0].Interface()
	})
}

func (this *Container) ProvideSingleton(provider interface{}) {
	providerType := reflect.TypeOf(provider)
	if providerType.Kind() != reflect.Func || providerType.NumOut() != 1 {
		panic(CallerTypeError)
	}
	resultType := providerType.Out(0)
	this.Singleton(utils.GetTypeKey(resultType), func() interface{} {
		return reflect.ValueOf(provider).Call(nil)[0].Interface()
	})
}

func New() Container {
	container := Container{}
	container.Flush()
	return container
}

func (this *Container) Bind(key string, provider contracts.InstanceProvider) {
	this.binds[this.GetKey(key)] = provider
}

func (this *Container) Instance(key string, instance interface{}) {
	this.instances[this.GetKey(key)] = instance
}

func (this *Container) Singleton(key string, provider contracts.InstanceProvider) {
	this.singletons[this.GetKey(key)] = provider
}

func (this *Container) Bound(key string) bool {
	key = this.GetKey(key)
	if _, existsBind := this.binds[key]; existsBind {
		return true
	}
	if _, existsSingleton := this.binds[key]; existsSingleton {
		return true
	}
	if _, existsInstance := this.binds[key]; existsInstance {
		return true
	}
	return false
}

func (this *Container) Alias(key string, alias string) {
	this.aliases[alias] = key
}

func (this *Container) GetKey(alias string) string {
	return utils.StringOr(this.aliases[alias], alias)
}

func (this *Container) Flush() {
	this.instances = make(map[string]interface{}, 0)
	this.singletons = make(map[string]contracts.InstanceProvider, 0)
	this.binds = make(map[string]contracts.InstanceProvider, 0)
	this.aliases = make(map[string]string, 0)
}

func (this *Container) Get(key string) interface{} {
	key = this.GetKey(key)
	if instance, existsInstance := this.instances[key]; existsInstance {
		return instance
	}
	if singletonProvider, existsProvider := this.singletons[key]; existsProvider {
		this.instances[key] = singletonProvider()
		return this.instances[key]
	}
	if instanceProvider, existsProvider := this.binds[key]; existsProvider {
		return instanceProvider()
	}
	return nil
}

func (this *Container) Call(fn interface{}, args ...interface{}) []interface{} {
	fnType := reflect.TypeOf(fn)
	argsTypeMap := NewArgumentsTypeMap(args)

	if fnType.Kind() != reflect.Func {
		panic(CallerTypeError)
	}

	argNum := fnType.NumIn()
	fnArgs := make([]reflect.Value, 0)

	for i := 0; i < argNum; i++ {
		arg := fnType.In(i)
		key := utils.GetTypeKey(arg)

		instance := utils.NotNil(argsTypeMap.Pull(key), func() interface{} {
			return this.Get(key)
		})

		if !reflect.TypeOf(instance).ConvertibleTo(arg) { // 不能转换，报错了，但是理论上来说不会走这里
			panic(errors.New(fmt.Sprintf("%s 的类型与预期不符", key)))
		}
		fnArgs = append(fnArgs, reflect.ValueOf(instance))
	}

	if argNum != len(fnArgs) {
		panic(ArgumentsLenError)
	}

	results := make([]interface{}, 0)

	for _, result := range reflect.ValueOf(fn).Call(fnArgs) {
		results = append(results, result.Interface())
	}

	return results
}
