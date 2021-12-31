package container

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"reflect"
	"sync"
)

var (
	CallerTypeError = errors.New("参数类型必须是有一个返回值的函数")
)

type Container struct {
	binds      map[string]contracts.MagicalFunc
	singletons map[string]contracts.MagicalFunc
	instances  sync.Map
	aliases    map[string]string
}

func newInstanceProvider(provider interface{}) contracts.MagicalFunc {
	if magicalFn := NewMagicalFunc(provider); magicalFn.NumOut() == 1 {
		return magicalFn
	}
	panic(CallerTypeError)
}

func New() contracts.Container {
	container := &Container{}
	container.Flush()
	return container
}

func (this *Container) Bind(key string, provider interface{}) {
	magicalFn := newInstanceProvider(provider)
	this.binds[this.GetKey(key)] = magicalFn
	this.Alias(key, utils.GetTypeKey(magicalFn.Returns()[0]))
}

func (this *Container) Instance(key string, instance interface{}) {
	this.instances.Store(this.GetKey(key), instance)
}

func (this *Container) Singleton(key string, provider interface{}) {
	magicalFn := newInstanceProvider(provider)
	this.singletons[this.GetKey(key)] = magicalFn
	this.Alias(key, utils.GetTypeKey(magicalFn.Returns()[0]))
}

func (this *Container) HasBound(key string) bool {
	key = this.GetKey(key)
	if _, existsBind := this.binds[key]; existsBind {
		return true
	}
	if _, existsSingleton := this.singletons[key]; existsSingleton {
		return true
	}
	if _, existsInstance := this.instances.Load(key); existsInstance {
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
	this.instances = sync.Map{}
	this.singletons = make(map[string]contracts.MagicalFunc, 0)
	this.binds = make(map[string]contracts.MagicalFunc, 0)
	this.aliases = make(map[string]string, 0)
}

func (this *Container) Get(key string, args ...interface{}) interface{} {
	key = this.GetKey(key)
	if tempInstance, existsInstance := this.instances.Load(key); existsInstance {
		return tempInstance
	}
	if singletonProvider, existsProvider := this.singletons[key]; existsProvider {
		value := this.Call(singletonProvider, args...)[0]
		this.instances.Store(key, value)
		return value
	}
	if instanceProvider, existsProvider := this.binds[key]; existsProvider {
		return this.Call(instanceProvider, args...)[0]
	}
	return nil
}

func (this *Container) Call(fn interface{}, args ...interface{}) []interface{} {
	magicalFn, isMagicalFunc := fn.(contracts.MagicalFunc)
	if !isMagicalFunc {
		magicalFn = NewMagicalFunc(fn)
	}
	argsTypeMap := NewArgumentsTypeMap(append(args, this))
	fnArgs := make([]reflect.Value, 0)

	for _, arg := range magicalFn.Arguments() {
		var (
			key      = utils.GetTypeKey(arg)
			argValue reflect.Value
		)

		tempInstance := argsTypeMap.Pull(key)

		if tempInstance == nil {
			tempInstance = this.Get(key, args...)
		}

		if tempInstance == nil {
			// 1. 尝试从外部参数注入
			tempInstance = argsTypeMap.FindConvertibleArg(arg)
			if tempInstance == nil {
				// 2. 尝试 new 一个自己然后 di 作为参数
				tempInstance = reflect.New(arg).Interface()
				this.DI(tempInstance, args...)
				argValue = reflect.ValueOf(tempInstance).Elem()
			} else {
				argValue = reflect.ValueOf(tempInstance).Convert(arg)
			}
		} else {
			argValue = reflect.ValueOf(tempInstance)
		}

		fnArgs = append(fnArgs, argValue)
	}

	results := make([]interface{}, 0)

	for _, result := range magicalFn.Call(fnArgs) {
		results = append(results, result.Interface())
	}

	return results
}

func (this *Container) DI(object interface{}, args ...interface{}) {
	fmt.Println(utils.GetTypeKey(reflect.TypeOf(object)))
	var (
		objectValue = reflect.ValueOf(object)
	)

	switch objectValue.Kind() {
	case reflect.Ptr:
		if objectValue.Elem().Kind() != reflect.Struct {
			panic(errors.New("参数必须是结构体指针!"))
		}
		objectValue = objectValue.Elem()
	default:
		panic(errors.New("参数必须是结构体指针!"))
	}

	var (
		valueType   = objectValue.Type()
		fieldNum    = objectValue.NumField()
		argsTypeMap = NewArgumentsTypeMap(append(args, this))
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
