package supports

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"reflect"
	"strings"
)

type Class struct {
	reflect.Type
}

type ClassException struct {
	error
	fields contracts.Fields
}

func (this ClassException) Error() string {
	return this.error.Error()
}

func (this ClassException) Fields() contracts.Fields {
	return this.fields
}

func GetClass(arg interface{}) Class {
	class := Class{reflect.TypeOf(arg)}
	if class.Kind() != reflect.Struct {
		panic(ClassException{
			errors.New("只支持 struct 类型!"),
			map[string]interface{}{
				"class": class.ClassName(),
			},
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
