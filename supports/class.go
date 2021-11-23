package supports

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/utils"
	"reflect"
	"strings"
)

type Class struct {
	reflect.Type
}

type ClassException struct {
	contracts.Exception
}

func GetClass(arg interface{}) Class {
	class := Class{reflect.TypeOf(arg)}
	if class.Kind() != reflect.Struct {
		panic(ClassException{
			Exception: exceptions.New("只支持 struct 类型!", map[string]interface{}{
				"class": class.ClassName(),
			}),
		})
	}
	return class
}

func (this *Class) New(data contracts.Fields) interface{} {
	object := reflect.New(this.Type)

	for name, value := range data {
		object.Elem().FieldByName(strings.Title(name)).Set(reflect.ValueOf(value))
	}

	return object.Elem().Interface()
}

func (this *Class) ClassName() string {
	return utils.GetTypeKey(this)
}
