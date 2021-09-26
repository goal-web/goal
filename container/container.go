package container

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"reflect"
)

var (
	CallerTypeError = errors.New("参数类型必须是有一个返回值的函数")
)

type Container struct {
	binds      map[string]contracts.InstanceProvider
	singletons map[string]contracts.InstanceProvider
	instances  map[string]interface{}
	aliases    map[string]string
}

func (this *Container) Provide(provider interface{}, aliases ...string) {
	providerType := reflect.TypeOf(provider)
	if providerType.Kind() != reflect.Func || providerType.NumOut() != 1 {
		panic(CallerTypeError)
	}

	resultType := providerType.Out(0)
	key := utils.GetTypeKey(resultType)

	this.Bind(key, func() interface{} {
		return reflect.ValueOf(provider).Call(nil)[0].Interface()
	})

	if alias := utils.StringOr(aliases...); alias != "" {
		this.Alias(key, alias)
	}
}

func (this *Container) ProvideSingleton(provider interface{}, aliases ...string) {
	providerType := reflect.TypeOf(provider)

	if providerType.Kind() != reflect.Func || providerType.NumOut() != 1 {
		panic(CallerTypeError)
	}

	resultType := providerType.Out(0)
	key := utils.GetTypeKey(resultType)

	this.Singleton(key, func() interface{} {
		return reflect.ValueOf(provider).Call(nil)[0].Interface()
	})

	if alias := utils.StringOr(aliases...); alias != "" {
		this.Alias(key, alias)
	}
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

func (this *Container) HasBound(key string) bool {
	key = this.GetKey(key)
	if _, existsBind := this.binds[key]; existsBind {
		return true
	}
	if _, existsSingleton := this.singletons[key]; existsSingleton {
		return true
	}
	if _, existsInstance := this.instances[key]; existsInstance {
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

		if instance == nil { // 不能注入 nil 参数
			instance = argsTypeMap.FindConvertibleArg(arg)
			if instance == nil {
				panic(errors.New(fmt.Sprintf("%s 无法注入 %s 参数", fnType.String(), arg.String())))
			}
		}

		fnArgs = append(fnArgs, reflect.ValueOf(instance))
	}

	results := make([]interface{}, 0)

	for _, result := range reflect.ValueOf(fn).Call(fnArgs) {
		results = append(results, result.Interface())
	}

	return results
}

func (this *Container) DI(object interface{}, args ...interface{}) {
	var (
		objectValue = reflect.ValueOf(object)
	)

	switch objectValue.Kind() {
	case reflect.Ptr:
		if objectValue.Elem().Kind() != reflect.Struct {
			panic(errors.New("参数必须是结构体指针 !"))
		}
		objectValue = objectValue.Elem()
	default:
		panic(errors.New("参数必须是结构体指针!"))
	}

	var (
		valueType   = objectValue.Type()
		fieldNum    = objectValue.NumField()
		argsTypeMap = NewArgumentsTypeMap(args)
		tempValue   = reflect.New(valueType).Elem()
		isComponent = valueType.Implements(ComponentType)
	)

	tempValue.Set(objectValue)

	for i := 0; i < fieldNum; i++ {
		var (
			field          = valueType.Field(i)
			key            = utils.GetTypeKey(field.Type)
			fieldTags      = utils.ParseStructTag(field.Tag)
			fieldValue     = tempValue.Field(i)
			fieldInterface interface{}
		)

		if di, existsDiTag := fieldTags["di"]; existsDiTag { // 配置了 fieldTags tag，优先用 tag 的配置
			fieldInterface = utils.NotNil(argsTypeMap.Pull(key), func() interface{} {
				if len(di) > 0 { // 如果指定某 di 值，优先取这个值
					return this.Get(di[0])
				}
				return nil
			}, func() interface{} {
				return this.Get(utils.StringOr(key))
			})
		} else if isComponent {
			fieldInterface = utils.NotNil(argsTypeMap.Pull(key), func() interface{} {
				return this.Get(key)
			})
		}

		if fieldInterface != nil {
			fieldType := reflect.TypeOf(fieldInterface)
			if fieldType.ConvertibleTo(field.Type) {
				fieldValue.Set(reflect.ValueOf(fieldInterface))
			} else {
				panic(errors.New(fmt.Sprintf("无法注入 %s ，因为类型不一致，目标类型为 %s，而将注入的类型为 %s", field.Name, field.Type.String(), fieldType.String())))
			}
		}
	}

	objectValue.Set(tempValue)

	return
}

func (this *Container) RegisterServices(services ...contracts.ServiceProvider) {
	for _, service := range services {
		service.Register(this)
	}
}
